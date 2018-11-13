package topic

import (
	"log"
	"net/http"
	//"regexp"

	"github.com/dangyanglim/go_cnode/mgoModels"
	//"github.com/dangyanglim/go_cnode/service/mail"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
)

var userModel = new(models.UserModel)
var topicModel = new(models.TopicModel)

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
	type Temp struct {
		Topic  models.Topic
		Author models.User
	}
	var temp Temp
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	id := c.Param("id")
	topic, author, _ := topicModel.GetTopicById(id)
	temp.Author = author
	temp.Topic = topic
	c.HTML(http.StatusOK, "topicIndex", gin.H{
		"title": "布局页面",
		"user":  user,
		"topic": temp,
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
