package models

import (
	//"log"

	db "github.com/dangyanglim/go_cnode/database"
	"gopkg.in/mgo.v2/bson"
	//"encoding/json"
	"time"
)

//"log"

type Reply struct {
	Id              bson.ObjectId `bson:"_id"`
	Topic_id           string        `json:"title"`
	Reply_id         string        `json:"content" `
	Author_id       bson.ObjectId `bson:"author_id" `
	Create_at       time.Time        `bson:"create_at"`
	Update_at       time.Time       `bson:"update_at"`
	Content		string `json:"content`
	Content_is_html bool          `json:"content_is_html"`

	Deleted         bool          `json:"deleted"`
}
type ReplyModel struct{}
func (p *ReplyModel) GetReplyById(id string) (reply Reply,err error) {
	mgodb := db.MogSession.DB("egg_cnode")
	objectId := bson.ObjectIdHex(id)

	err = mgodb.C("replies").Find(bson.M{"_id": objectId}).One(&reply)


	return reply, err
}






