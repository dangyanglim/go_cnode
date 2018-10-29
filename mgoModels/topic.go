package models

import (
	//"log"
	"gopkg.in/mgo.v2/bson"
	db "github.com/dangyanglim/go_cnode/database"
)

//"log"

type Topic struct {
	Title        string    `json:"title"`
	Content string `json:"content" `
	Author_id    string `json:"author_id" `
	Top bool `json:"top" `
	Good bool `json:"good" `
	Lock    bool `json:"lock"`
	Reply_count	uint `json:"reply_count"`
	visit_count	uint `json:"visit_count"`
	Collect_count	uint `json:"collect_count"`
	Create_at	string `json:"create_at"`
	Update_at string `json:"update_at"`
	Last_reply uint `json:"last_reply"`
	Last_reply_at string `json:"last_reply_at"`
	Content_is_html bool `json:"content_is_html"`
	Tab string `json:"tab"`
	deleted bool `json:"deleted"`
}
type TopicModel struct{}
func (p *TopicModel) GetTopicByQuery(tab string,good bool)(topics []Topic,err error){
	mgodb := db.MogSession.DB("egg_cnode")
	if tab==""||tab=="all"{
		err=mgodb.C("topics").Find(bson.M{"good":good}).All(&topics);
	}else{
		err=mgodb.C("topics").Find(bson.M{"tab":tab,"good":good}).All(&topics);
	}
	
	return topics,err 	
}


