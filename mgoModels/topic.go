package models

import (
	//"log"
	"encoding/json"
	db "go_cnode/database"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

//"log"

type Topic struct {
	Id               bson.ObjectId `bson:"_id"`
	Title            string        `json:"title"`
	Content          string        `json:"content" `
	Author_id        bson.ObjectId `bson:"author_id" `
	Top              bool          `json:"top" `
	Good             bool          `json:"good" `
	Lock             bool          `json:"lock"`
	Reply_count      uint          `json:"reply_count"`
	Visit_count      uint          `json:"visit_count"`
	Collect_count    uint          `json:"collect_count"`
	Create_at        time.Time     `bson:"create_at"`
	Create_at_string string        `json:"create_at_string,omitempty"`
	Update_at        string        `json:"update_at"`
	Last_reply       bson.ObjectId `bson:"last_reply,omitempty"`
	Last_reply_at    time.Time     `json:"last_reply_at,omitempty"`
	Content_is_html  bool          `json:"content_is_html"`
	Tab              string        `json:"tab"`
	Deleted          bool          `json:"deleted"`
}
type TopicModel struct{}

var userModel = new(UserModel)
var replyModel = new(ReplyModel)
var topicModel = new(TopicModel)

type TopciAndAuthor struct {
	Author User  `json:"author"`
	Topic  Topic `json:"topic"`
	Reply  Reply `json:"reply"`
}

func (p *TopicModel) GetTopicByQuery(tab string, good bool, limit int, skip int) (topics []Topic, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	if tab == "" || tab == "all" {
		err = mgodb.C("topics").Find(bson.M{"good": good,"deleted":false}).Sort("-top", "-create_at").Limit(limit).Skip(skip).All(&topics)
	} else {
		err = mgodb.C("topics").Find(bson.M{"tab": tab, "good": good,"deleted":false}).Sort("-top", "-create_at").Limit(limit).Skip(skip).All(&topics)
	}
	//log.Println(topics)
	return topics, err
}
func (p *TopicModel) GetTopicBy(tab string, good bool, limit int, skip int) (topics []Topic, topicss []byte, err error) {

	var temps []TopciAndAuthor
	mgodb := db.MogSession.DB("egg_cnode")
	if tab == "" || tab == "all" {
		err = mgodb.C("topics").Find(bson.M{"good": good,"deleted":false}).Sort("-top", "-create_at").Limit(limit).Skip(skip).All(&topics)
	} else {
		err = mgodb.C("topics").Find(bson.M{"tab": tab, "good": good,"deleted":false}).Sort("-top", "-create_at").Limit(limit).Skip(skip).All(&topics)
	}
	//log.Println(topics)
	for _, v := range topics {
		var temp TopciAndAuthor
		temp.Topic = v
		author, _ := userModel.GetUserById(v.Author_id.Hex())
		temp.Author = author
		if v.Last_reply.Hex() != "" {
			reply, _ := replyModel.GetReplyById(v.Last_reply.Hex())
			temp.Reply = reply
		}

		temps = append(temps, temp)
	}
	topicss, _ = json.Marshal(temps)

	return topics, topicss, err
}
func (p *TopicModel) GetTopicByQueryCount(tab string, good bool) (count int, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	if tab == "" || tab == "all" {
		count, err = mgodb.C("topics").Find(bson.M{"deleted":false}).Count()
	} else {
		if good == true {
			count, err = mgodb.C("topics").Find(bson.M{"good": good,"deleted":false}).Count()
		} else {
			count, err = mgodb.C("topics").Find(bson.M{"tab": tab,"deleted":false}).Count()
		}

	}

	return count, err
}
func (p *TopicModel) GetTopicByAuthorQueryCount(objectId bson.ObjectId, tab string, good bool) (count int, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	if tab == "" || tab == "all" {
		count, err = mgodb.C("topics").Find(bson.M{"author_id": objectId,"deleted":false}).Count()
	} else {
		if good == true {
			count, err = mgodb.C("topics").Find(bson.M{"author_id": objectId, "good": good,"deleted":false}).Count()
		} else {
			count, err = mgodb.C("topics").Find(bson.M{"author_id": objectId, "tab": tab,"deleted":false}).Count()
		}

	}

	return count, err
}

// type ReplyAndAuthor struct {
// 	Author models.User
// 	Reply  models.Reply
// }
func (p *TopicModel) GetTopicById(id string) (topic Topic, err error) {
	//mgodb := db.MogSession.DB("egg_cnode")
	mgodb:=db.Mgodb
	objectId := bson.ObjectIdHex(id)
	err = mgodb.C("topics").Find(bson.M{"_id": objectId}).One(&topic)
	return topic, err
}
func (p *TopicModel) GetTopicByIdWithReply(id string) (topic Topic, author User, replies []Reply, repliyWithAuthors []ReplyAndAuthor, err error) {
	mgodb:=db.Mgodb
	objectId := bson.ObjectIdHex(id)

	err = mgodb.C("topics").Find(bson.M{"_id": objectId}).One(&topic)

	author, _ = userModel.GetUserById(topic.Author_id.Hex())
	topic.Create_at_string = topic.Create_at.Format("2006-01-02 15:04:05")
	replies, repliyWithAuthors, _ = replyModel.GetRepliesByTopicId(topic.Id.Hex())
	return topic, author, replies, repliyWithAuthors, err
}
func (p *TopicModel) NewAndSave(title string, tab string, id string, content string) (topic Topic, err error) {

	objectId := bson.ObjectIdHex(id)
	topic = Topic{
		Id:        bson.NewObjectId(),
		Title:     title,
		Content:   content,
		Tab:       tab,
		Author_id: objectId,
		Create_at: time.Now(),
	}
	mgodb:=db.Mgodb
	err = mgodb.C("topics").Insert(&topic)
	log.Println(err)
	return topic, err
}
func (p *TopicModel) GetTopicNoReply() (topics []Topic, err error) {
	mgodb:=db.Mgodb

	err = mgodb.C("topics").Find(bson.M{"reply_count": 0,"deleted":false}).Sort("-create_at").Limit(5).All(&topics)

	return topics, err
}
func (p *TopicModel) GetAuthorOtherTopics(author_id string, topic_id string) (topics []Topic, err error) {
	mgodb:=db.Mgodb
	objectId := bson.ObjectIdHex(author_id)
	topic_objectId := bson.ObjectIdHex(topic_id)
	err = mgodb.C("topics").Find(bson.M{"author_id": objectId,"deleted":false, "_id": bson.M{"$nin": []bson.ObjectId{topic_objectId}}}).Limit(5).Sort("-last_reply_at").All(&topics)
	return topics, err
}
func (p *TopicModel) GetAuthorTopics(author_id string, limit int, skip int) (topics []Topic, topicss []byte, err error) {
	var temps []TopciAndAuthor
	mgodb:=db.Mgodb
	objectId := bson.ObjectIdHex(author_id)
	err = mgodb.C("topics").Find(bson.M{"author_id": objectId,"deleted":false}).Skip(skip).Limit(limit).Sort("-create_at").All(&topics)
	for _, v := range topics {
		var temp TopciAndAuthor
		temp.Topic = v
		author, _ := userModel.GetUserById(v.Author_id.Hex())
		temp.Author = author
		if v.Last_reply.Hex() != "" {
			reply, _ := replyModel.GetReplyById(v.Last_reply.Hex())
			temp.Reply = reply
		}

		temps = append(temps, temp)
	}
	topicss, _ = json.Marshal(temps)
	return topics, topicss, err
}
func (p *TopicModel) GetReplyTopics(author_id string, limit int, skip int, most int) (topicss []byte, err error) {
	var temps []TopciAndAuthor
	var replies []Reply
	var topic_ids map[bson.ObjectId]int
	var topic_id_ints []bson.ObjectId
	topic_ids = make(map[bson.ObjectId]int)
	mgodb:=db.Mgodb
	objectId := bson.ObjectIdHex(author_id)
	err = mgodb.C("replies").Find(bson.M{"author_id": objectId,"deleted":false}).Sort("-create_at").Skip(skip).Limit(limit).All(&replies)
	for i := 0; i < len(replies); i++ {
		if _, ok := topic_ids[replies[i].Topic_id]; ok {
		} else {
			topic_ids[replies[i].Topic_id] = len(topic_ids) + 1
			topic_id_ints = append(topic_id_ints, replies[i].Topic_id)
		}
		if len(topic_ids) == most {
			break
		}
	}
	for j := 0; j < len(topic_id_ints); j++ {
		topic, err2 := topicModel.GetTopicById(topic_id_ints[j].Hex())
		if err2 == nil && topic.Author_id.Hex() != "" {
			var temp TopciAndAuthor
			temp.Topic = topic
			author, _ := userModel.GetUserById(topic.Author_id.Hex())
			temp.Author = author
			if topic.Last_reply.Hex() != "" {
				reply, _ := replyModel.GetReplyById(topic.Last_reply.Hex())
				temp.Reply = reply
			}
			temps = append(temps, temp)
		}
	}
	topicss, _ = json.Marshal(temps)
	return topicss, err
}
func (p *TopicModel) UpdateReplyCount(id string, replyId bson.ObjectId) (err error) {
	mgodb:=db.Mgodb
	objectId := bson.ObjectIdHex(id)
	//objectReplyId := bson.ObjectIdHex(replyId)
	err = mgodb.C("topics").Update(bson.M{"_id": objectId},
		bson.M{
			"$inc": bson.M{"reply_count": 1},
			"$set": bson.M{"last_reply_at": time.Now(), "last_reply": replyId},
		})
	log.Println(err)
	return err
}
func (p *TopicModel) UpdateVisitCount(id string) (err error) {
	mgodb:=db.Mgodb
	objectId := bson.ObjectIdHex(id)

	err = mgodb.C("topics").Update(bson.M{"_id": objectId},
		bson.M{
			"$inc": bson.M{"visit_count": 1},
		})
	log.Println(err)
	return err
}
func (p *TopicModel) SetTop(id string, value bool) (err error) {
	mgodb:=db.Mgodb
	objectId := bson.ObjectIdHex(id)

	err = mgodb.C("topics").Update(bson.M{"_id": objectId},
		bson.M{
			"$set": bson.M{"top": value},
		})
	log.Println(err)
	return err
}
func (p *TopicModel) Delete( id string) (err error) {

	objectId := bson.ObjectIdHex(id)
	mgodb:=db.Mgodb
	err = mgodb.C("topics").Update(bson.M{"_id": objectId},
		bson.M{
			"$set": bson.M{"update_at": time.Now(), "deleted": true},
		})
	return err
}
func (p *TopicModel) Update(topic Topic) (err error) {
	mgodb:=db.Mgodb
	err = mgodb.C("topics").Update(bson.M{"_id": topic.Id},
		bson.M{
			"$set": bson.M{"update_at": time.Now(), "title": topic.Title,
			
					"tab":topic.Tab,"content":topic.Content},
		})
	log.Println(err)
	return err
}
