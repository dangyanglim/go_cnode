package models

import (
	//"log"
	"log"

	db "github.com/dangyanglim/go_cnode/database"
	"gopkg.in/mgo.v2/bson"
)

//"log"

type Topic struct {
	Id              bson.ObjectId `bson:"_id"`
	Title           string        `json:"title"`
	Content         string        `json:"content" `
	Author_id       bson.ObjectId `bson:"author_id" `
	Top             bool          `json:"top" `
	Good            bool          `json:"good" `
	Lock            bool          `json:"lock"`
	Reply_count     uint          `json:"reply_count"`
	Visit_count     uint          `json:"visit_count"`
	Collect_count   uint          `json:"collect_count"`
	Create_at       string        `json:"create_at"`
	Update_at       string        `json:"update_at"`
	Last_reply      uint          `json:"last_reply"`
	Last_reply_at   string        `json:"last_reply_at"`
	Content_is_html bool          `json:"content_is_html"`
	Tab             string        `json:"tab"`
	Deleted         bool          `json:"deleted"`
}
type TopicModel struct{}

var userModel = new(UserModel)

func (p *TopicModel) GetTopicByQuery(tab string, good bool) (topics []Topic, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	if tab == "" || tab == "all" {
		err = mgodb.C("topics").Find(bson.M{"good": good}).All(&topics)
	} else {
		err = mgodb.C("topics").Find(bson.M{"tab": tab, "good": good}).All(&topics)
	}

	return topics, err
}
func (p *TopicModel) GetTopicById(id string) (topic Topic, author User, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	objectId := bson.ObjectIdHex(id)

	err = mgodb.C("topics").Find(bson.M{"_id": objectId}).One(&topic)

	author, _ = userModel.GetUserById(topic.Author_id.Hex())
	log.Println(topic)
	log.Println(author)

	return topic, author, err
}
func (p *TopicModel) NewAndSave(title string, tab string, id string, content string) ( topic Topic,err error) {

	
	objectId := bson.ObjectIdHex(id)
	topic = Topic{
		Id:          bson.NewObjectId(),
		Title:        title,
		Content:   content,
		Tab:        tab,
		Author_id:objectId,
	}
	mgodb := db.MogSession.DB("egg_cnode")
	err = mgodb.C("topics").Insert(&topic)
	log.Println(err)
	return topic,err
}
func (p *TopicModel) GetTopicNoReply() (topics []Topic, err error) {
	mgodb := db.MogSession.DB("egg_cnode")

	err = mgodb.C("topics").Find(bson.M{"reply_count": 0}).Limit(5).All(&topics)
	

	return topics, err
}
