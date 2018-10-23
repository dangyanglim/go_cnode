package site

import (
	"net/http"

	//. "github.com/dangyanglim/go_cnode/models"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	//c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "布局页面",
		"aa":    "aa",
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
