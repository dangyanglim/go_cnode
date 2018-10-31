package topic

import (
	//"log"
	"net/http"
	//"regexp"

	"github.com/dangyanglim/go_cnode/mgoModels"
	//"github.com/dangyanglim/go_cnode/service/mail"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
)

var userModel = new(models.UserModel)
var topicModel = new(models.TopicModel)







func TopicCreate(c *gin.Context){
	session := sessions.Get(c)
	var name string
	user := models.User{}
	//var err error

	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	tabs:=[3]map[string]string{{"value":"share","text":"分享"},{"value":"ask", "text":"问答"},{"value":"job","text":"招聘"}}	
	c.HTML(http.StatusOK, "edit", gin.H{
		"user":user,
		"tabs":tabs,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},    
	})	
}

func Index(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}	
	id := c.Param("id")
	topic,_:=topicModel.GetTopicById(id)
	c.HTML(http.StatusOK, "topicIndex", gin.H{
		"title": "布局页面",
		"user":user,
		"topic":topic,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}