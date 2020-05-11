package models

import (
	db "go_cnode/database"
	"gopkg.in/mgo.v2/bson"
	"time"
	"log"
)

//"log"

type Message struct {
	Id        bson.ObjectId `bson:"_id"`
	Type      string        `json:"type"`
	Master_id bson.ObjectId `json:"master_id" `
	Author_id bson.ObjectId `bson:"author_id" `
	Topic_id  bson.ObjectId `json:"topic_id" `
	Reply_id  bson.ObjectId `json:"reply_id" `
	Has_read  bool          `json:"has_read"`
	Create_at time.Time     `bson:"create_at"`
}
type MessageModel struct{}

func (p *MessageModel) GetMessagesCount(id string) (count int, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	objectId := bson.ObjectIdHex(id)
	count, err = mgodb.C("messages").Find(bson.M{"master_id": objectId,"has_read": false}).Count()
	return count, err
}
func (p *MessageModel) GetMessageById(id string) (message Message, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	objectId := bson.ObjectIdHex(id)
	err = mgodb.C("messages").Find(bson.M{"_id": objectId}).One(&message)
	return message, err
}
func (p *MessageModel) GetMessagesByUserId(id bson.ObjectId) (messages []Message, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	
	err = mgodb.C("messages").Find(bson.M{"master_id": id, "has_read": true}).Sort("-create_at").All(&messages)
	return messages, err
}
func (p *MessageModel) GetUnreadMessagesByUserId(id bson.ObjectId) (messages []Message, err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	err = mgodb.C("messages").Find(bson.M{"master_id": id, "has_read": false}).Sort("-create_at").All(&messages)
	return messages, err
}
func (p *MessageModel) UpdateOneMessageToRead(msgId string) (err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	objectId := bson.ObjectIdHex(msgId)
	err = mgodb.C("messages").Update(bson.M{"_id": objectId},
		bson.M{
			"$set": bson.M{"has_read": true},
		})
	return err
}
func (p *MessageModel) UpdateMessagesToRead(userId string, messages []Message) (err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	var ids []bson.ObjectId
	for _, message := range messages {
		ids = append(ids, message.Id)
	}
	log.Println(ids)
	_,err = mgodb.C("messages").UpdateAll(
		bson.M{"master_id": bson.ObjectIdHex(userId),
			"_id":bson.M{"$in": ids}},
		bson.M{
			"$set": bson.M{"has_read": true},
		})
	return err
}
func (p *MessageModel) SendAtMessage(userId string, authorId string, topicId string, replyId bson.ObjectId) (err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	message := Message{
		Id:        bson.NewObjectId(),
		Type:      "at",
		Master_id: bson.ObjectIdHex(userId),
		Topic_id:  bson.ObjectIdHex(topicId),
		Author_id: bson.ObjectIdHex(authorId),
		Reply_id:  replyId,
		Create_at: time.Now(),
	}
	err = mgodb.C("messages").Insert(&message)
	return err
}
func (p *MessageModel) SendReplyMessage(userId string, authorId string, topicId string, replyId bson.ObjectId) (err error) {
	//mgodb := db.MogSession.DB("egg_cnode")
	mgodb:=db.Mgodb
	message := Message{
		Id:        bson.NewObjectId(),
		Type:      "reply",
		Master_id: bson.ObjectIdHex(userId),
		Topic_id:  bson.ObjectIdHex(topicId),
		Author_id: bson.ObjectIdHex(authorId),
		Reply_id:  replyId,
		Create_at: time.Now(),
	}
	err = mgodb.C("messages").Insert(&message)
	return err
}
