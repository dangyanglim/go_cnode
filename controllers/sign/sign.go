package sign

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/dangyanglim/go_cnode/mgoModels"

)

var userModel = new(models.UserModel)
func Signup(c *gin.Context) {
	
	c.HTML(http.StatusOK, "signup", gin.H{
		"title": "布局页面",
		"config":gin.H{
			"description":"CNode：Node.js专业中文社区",
		},		
	})
}
func Signin(c *gin.Context) {
	
	c.HTML(http.StatusOK, "signin", gin.H{
		"title": "布局页面",
		"config":gin.H{
			"description":"CNode：Node.js专业中文社区",
		},
	})
}
func Login(c *gin.Context) {
	name := c.Request.FormValue("name")
	//pass := c.Request.FormValue("pass")
	user,_:=userModel.GetUserByName(name)
	log.Println(user)
	var no_reply_topics =[]string{"2","2"};
	var tops =[]string{"2","2"};
	tab:=c.Request.FormValue("tab")
	
	if tab==""{
	  tab="all"
	}	
	
	//c.Redirect(http.StatusMovedPermanently, "http://shanshanpt.github.io/")
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "布局页面",
	"no_reply_topics":no_reply_topics,
	"user":user,
    "tops":tops,
    "tab":tab,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
	//c.Redirect(301,"/")	
}
func SearchPass(c *gin.Context) {
	
	c.HTML(http.StatusOK, "search_pass", gin.H{
		"title": "布局页面",
		"config":gin.H{
			"description":"CNode：Node.js专业中文社区",
		},
	})
}



