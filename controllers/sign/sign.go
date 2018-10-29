package sign

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

func ShowSignup(c *gin.Context) {

	c.HTML(http.StatusOK, "showsignup", gin.H{
		"title": "布局页面",
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
func Signin(c *gin.Context) {

	c.HTML(http.StatusOK, "signin", gin.H{
		"title": "布局页面",
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
func Signout(c *gin.Context) {
	session := sessions.Get(c)
	session.Clear()
	session.Save()
	c.Redirect(301, "/")
}

func Login(c *gin.Context) {
	name := c.Request.FormValue("name")
	//pass := c.Request.FormValue("pass")
	user, _ := userModel.GetUserByName(name)
	session := sessions.Get(c)
	session.Set("loginname", user.Loginname)
	session.Set("accessToken", user.AccessToken)
	session.Save()
	log.Println(user.Loginname)
	// var no_reply_topics =[]string{"2","2"};
	// var tops =[]string{"2","2"};
	tab := c.Request.FormValue("tab")

	if tab == "" {
		tab = "all"
	}

	//c.Redirect(http.StatusMovedPermanently, "http://shanshanpt.github.io/")
	// c.HTML(http.StatusOK, "index", gin.H{
	// 	"title": "布局页面",
	// "no_reply_topics":no_reply_topics,
	// "user":user,
	// "tops":tops,
	// "tab":tab,
	// 	"config": gin.H{
	// 		"description": "CNode：Node.js专业中文社区",
	// 	},
	// })
	c.Redirect(301, "/")
}
func SearchPass(c *gin.Context) {

	c.HTML(http.StatusOK, "search_pass", gin.H{
		"title": "布局页面",
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
func Setting(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	//var err error

	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	log.Println(user)
	c.HTML(http.StatusOK, "setting", gin.H{
		"title": "布局页面",
		"user":  user,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
func Message(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	//var err error

	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	log.Println(user)
	c.HTML(http.StatusOK, "message_index", gin.H{
		"title": "布局页面",
		"user":  user,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
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
	user, err := userModel.GetUserByNameOrEmail(loginname, email)
	log.Println(user)
	log.Println(err)
	if err == nil {
		c.HTML(http.StatusOK, "signup", gin.H{
			"title": "布局页面",
			// "user":  user,
			"error":     "用户名或邮箱已被使用",
			"loginname": loginname,
			"email":     email,
			"config": gin.H{
				"description": "CNode：Node.js专业中文社区",
			},
		})
		return
	}
	userModel.NewAndSave(loginname, loginname, email, pass, "aa", true)
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
