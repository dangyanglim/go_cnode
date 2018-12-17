package models

import (
	"log"

	db "github.com/dangyanglim/go_cnode/database"
	"gopkg.in/mgo.v2/bson"
	//"encoding/json"
	"time"
)

//"log"

type Reply struct {
	Id              bson.ObjectId `bson:"_id"`
	Topic_id           bson.ObjectId        `bson:"topic_id"`
	Reply_id         string        `json:"content" `
	Author_id       bson.ObjectId `bson:"author_id" `
	Create_at       time.Time        `bson:"create_at"`
	Update_at       time.Time       `bson:"update_at"`
	Content		string `json:"content`
	Content_is_html bool          `json:"content_is_html"`

	Deleted         bool          `json:"deleted"`
}
type ReplyAndAuthor struct {
	Reply Reply
	Author User
	
}
type ReplyModel struct{}
func (p *ReplyModel) GetReplyById(id string) (reply Reply,err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	objectId := bson.ObjectIdHex(id)
	err = mgodb.C("replies").Find(bson.M{"_id": objectId}).One(&reply)
	return reply, err
}
func (p *ReplyModel) GetRepliesByTopicId(id string) (replies []Reply,replyAndAuthor []ReplyAndAuthor,err error) {
	//var replyAndAuthor []ReplyAndAuthor
	mgodb := db.MogSession.DB("egg_cnode")
	objectId := bson.ObjectIdHex(id)
	err = mgodb.C("replies").Find(bson.M{"topic_id": objectId}).All(&replies)
	for _,v:=range replies{
		var temp ReplyAndAuthor
		log.Println("aa")
		author, _ := userModel.GetUserById(v.Author_id.Hex())
		temp.Reply=v
		temp.Author=author
		replyAndAuthor=append(replyAndAuthor,temp)
		log.Println(temp)
	}
	return replies,replyAndAuthor, err
}






