package reply

import (
	"log"
	"net/http"
	//"regexp"

	"github.com/dangyanglim/go_cnode/mgoModels"
	//"github.com/dangyanglim/go_cnode/service/mail"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"github.com/dangyanglim/go_cnode/service/cache"
	"encoding/json"
)

var userModel = new(models.UserModel)
var topicModel = new(models.TopicModel)
var replyModel = new(models.ReplyModel)
func ShowCreate(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	//var err error

	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	tabs := [3]map[string]string{{"value": "share", "text": "分享"}, {"value": "ask", "text": "问答"}, {"value": "job", "text": "招聘"}}
	c.HTML(http.StatusOK, "edit", gin.H{
		"user": user,
		"tabs": tabs,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}

func Index(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	var no_reply_topics []models.Topic;
	type Temp struct {
		Topic  models.Topic
		Author models.User
		Replies []models.Reply
		RepliyWithAuthors []models.ReplyAndAuthor
	}
	var temp Temp
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	id := c.Param("id")
	topic, author,replies, repliyWithAuthors,_:= topicModel.GetTopicById(id)
	temp.Author = author
	temp.Topic = topic
	temp.Replies=replies
	NoOfRepliy:=len(replies)
	temp.RepliyWithAuthors=repliyWithAuthors
	no_reply_topics2,err2:=cache.Get("no_reply_topics")
	json.Unmarshal(no_reply_topics2.([]byte),&no_reply_topics)
	log.Println("temp")
	log.Println(err2)
	//log.Println(temp)
	if(err2!=nil){
	  no_reply_topics,_=topicModel.GetTopicNoReply()
	  no_reply_topics_json,_:=json.Marshal(no_reply_topics)
	  cache.SetEx("no_reply_topics",no_reply_topics_json)
	}
	other_topics,_:=topicModel.GetAuthorOtherTopics(author.Id.Hex(),id)
	c.HTML(http.StatusOK, "topicIndex", gin.H{
		"title": "布局页面",
		"user":  user,
		"topic": temp,
		"NoOfRepliy":NoOfRepliy,
		"no_reply_topics":no_reply_topics,
		"author_other_topics": other_topics,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
func Create(c *gin.Context) {
	session := sessions.Get(c)
	var name string	
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	log.Println(user)
	id:=user.Id.Hex()
	log.Println(id)	
	tab := c.Request.FormValue("tab")
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	topic,_:=topicModel.NewAndSave(title,tab,id,content)
	url:="/topic/"+topic.Id.Hex()
	c.Redirect(301, url)
}
func Add(c *gin.Context) {
	topic_id:=c.Param("topic_id")
	r_content := c.Request.FormValue("r_content")
	user_id := c.Request.FormValue("user_id")
	log.Println(topic_id)
	log.Println(r_content)
	topicModel.UpdateReplyCount(topic_id)
	replyModel.NewAndSave(r_content,topic_id,user_id,"")
	url:="/topic/"+topic_id
	c.Redirect(301, url)	
}
