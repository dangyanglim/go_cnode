package sign

import (
	"net/http"

	"github.com/gin-gonic/gin"


)


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
func SearchPass(c *gin.Context) {
	
	c.HTML(http.StatusOK, "search_pass", gin.H{
		"title": "布局页面",
		"config":gin.H{
			"description":"CNode：Node.js专业中文社区",
		},
	})
}



