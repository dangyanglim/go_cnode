package topic

import (
	"log"
	"net/http"
	"regexp"

	"github.com/dangyanglim/go_cnode/mgoModels"
	"github.com/dangyanglim/go_cnode/service/mail"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
)

var userModel = new(models.UserModel)







func Signup(c *gin.Context) {
	var msg string
	/* 	session := sessions.Get(c)
	   	var name string
	   	user := models.User{}
	   	//var err error

	   	if nil != session.Get("loginname") {
	   		name = session.Get("loginname").(string)
	   		user, _ = userModel.GetUserByName(name)
	   	}
		   log.Println(user) */
	loginname := c.Request.FormValue("loginname")
	email := c.Request.FormValue("email")
	pass := c.Request.FormValue("pass")
	rePass := c.Request.FormValue("re_pass")
	pat := `^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`
	reg, _ := regexp.Match(pat, []byte(email))
	if loginname == "" || email == "" || pass == "" || rePass == "" {
		msg = "信息不完整。"
	} else if len(loginname) < 5 {
		msg = "用户名至少需要5个字符。"
	} else if rePass != pass {
		msg = "两次密码输入不一致。"
	} else if !reg {
		msg = "邮箱不合法。"
	}

	log.Print(reg)
	if msg != "" {
		c.HTML(http.StatusOK, "signup", gin.H{
			"title": "布局页面",
			// "user":  user,
			"error":     msg,
			"loginname": loginname,
			"email":     email,
			"config": gin.H{
				"description": "CNode：Node.js专业中文社区",
			},
		})
		return
	}
	mail.SendActiveMail(email, "aaa", loginname)
	c.HTML(http.StatusOK, "signup", gin.H{
		"title": "布局页面",
		// "user":  user,
		"success":   "欢迎加入 ！我们已给您的注册邮箱发送了一封邮件，请点击里面的链接来激活您的帐号。",
		"loginname": loginname,
		"email":     email,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})

}
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
