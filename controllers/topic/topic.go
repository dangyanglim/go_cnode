package topic

import (
	"log"
	"net/http"
	//"regexp"
	//"fmt"
	//"os"
	//"io"
	"go_cnode/mgoModels"
	//"github.com/dangyanglim/go_cnode/service/mail"
	"encoding/json"
	"go_cnode/service/cache"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"github.com/russross/blackfriday"
	"html/template"/*  */
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
	var no_reply_topics []models.Topic
	type Temp struct {
		Topic             models.Topic
		LinkContent 		template.HTML
		Author            models.User
		Replies           []models.Reply
		RepliyWithAuthors []models.ReplyAndAuthor
	}
	var temp Temp
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	id := c.Param("id")
	topic, author, replies, repliyWithAuthors, _ := topicModel.GetTopicById(id)
	temp.Author = author
	temp.LinkContent= template.HTML(blackfriday.Run([]byte(topic.Content)))
	temp.Topic = topic
	temp.Replies = replies
	NoOfRepliy := len(replies)
	temp.RepliyWithAuthors = repliyWithAuthors
	no_reply_topics2, err2 := cache.Get("no_reply_topics")
	json.Unmarshal(no_reply_topics2.([]byte), &no_reply_topics)
	
	if err2 != nil {
		no_reply_topics, _ = topicModel.GetTopicNoReply()
		no_reply_topics_json, _ := json.Marshal(no_reply_topics)
		cache.SetEx("no_reply_topics", no_reply_topics_json)
	}
	log.Println(temp)
	topicModel.UpdateVisitCount(id)
	other_topics, _ := topicModel.GetAuthorOtherTopics(author.Id.Hex(), id)
	c.HTML(http.StatusOK, "topicIndex", gin.H{
		"title":               "布局页面",
		"user":                user,
		"topic":               temp,
		"NoOfRepliy":          NoOfRepliy,
		"no_reply_topics":     no_reply_topics,
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
	id := user.Id.Hex()
	log.Println(id)
	tab := c.Request.FormValue("tab")
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	topic, _ := topicModel.NewAndSave(title, tab, id, content)
	url := "/topic/" + topic.Id.Hex()
	c.Redirect(301, url)
}
const (
	upload_path string = "./public/upload/"
	upload_path2 string = "/public/upload/"
)
func Upload(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	log.Println(user)
	id := user.Id.Hex()
	log.Println(id)
	picName := c.Request.FormValue("name")
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	log.Println(picName)

	//创建文件
	// fW, err := os.Create(upload_path + file.Filename)
	// if err != nil {
	// 	fmt.Println("文件创建失败")
	// 	return
	// }
	// defer fW.Close()
	// _, err = io.Copy(fW, file)
	// if err != nil {
	// 	fmt.Println("文件保存失败")
	// 	return
	// }
	c.SaveUploadedFile(file,upload_path + file.Filename)
	var msg struct {
		Success    bool `json:"success"`
		Url string `json:"url"`
	}
	msg.Success = true
	msg.Url=upload_path2 + file.Filename	
	c.JSON(http.StatusOK, msg)


}
