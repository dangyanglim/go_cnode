package main

import (
	"net/http"

	"github.com/dangyanglim/go_cnode/controllers/sign"
	"github.com/dangyanglim/go_cnode/controllers/site"
	"github.com/dangyanglim/go_cnode/controllers/topic"
	"github.com/dangyanglim/go_cnode/controllers/reply"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	store := sessions.NewCookieStore([]byte("secret123"))
	router.Use(sessions.Middleware("my_session", store))
	router.LoadHTMLGlob("views/*")
	router.StaticFS("/public", http.Dir("./public"))
	router.StaticFile("/favicon.ico", "./public/images/cnode_icon_32.png")
	router.GET("/", site.Index)
	router.GET("/about", site.About)
	router.GET("/signup", sign.ShowSignup)
	router.POST("/signup", sign.Signup)
	router.POST("/signout", sign.Signout)
	router.GET("/signin", sign.Signin)
	router.GET("/setting", sign.Setting)
	router.GET("/my/messages", sign.Message)

	router.POST("/passport/local", sign.Login)
	router.GET("/search_pass", sign.SearchPass)
	router.GET("/api", site.Api)
	router.GET("/getstart", site.Getstart)
	router.GET("/topic/:id", topic.Index)
	router.GET("/topics/create", topic.ShowCreate)
	router.POST("/topic/create", topic.Create)
	router.GET("/active_account", sign.ActiveAccount) // 帐号激活
	router.POST("/reply/:topic_id",reply.Add)
	return router
}
