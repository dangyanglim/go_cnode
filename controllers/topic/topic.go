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
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"github.com/tommy351/gin-sessions"
	"go_cnode/service/cache"
	"html/template" /*  */
	"strings"
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
		LinkContent       template.HTML
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
	topic, author, replies, repliyWithAuthors, _ := topicModel.GetTopicByIdWithReply(id)
	temp.Author = author
	topic.Content = strings.Replace(topic.Content, "\r\n", "<br/>", -1)
	//temp.LinkContent=topic.Content
	temp.LinkContent = template.HTML(blackfriday.Run([]byte(topic.Content)))
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
func Top(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	if user.Name != "admin" {
		c.HTML(http.StatusOK, "notify", gin.H{
			"error": "没权限",
		})
		return
	}
	id := c.Param("id")
	topic, err := topicModel.GetTopicById(id)
	if err != nil {
		c.HTML(http.StatusOK, "notify", gin.H{
			"error": "话题不存在",
		})
		return
	}
	topicModel.SetTop(id, !topic.Top)
	msg := ""
	if topic.Top {
		msg = "此话题取消置顶成功。"
	} else {
		msg = "此话题置顶成功。"
	}

	c.HTML(http.StatusOK, "notify", gin.H{
		"success": msg,
	})
}
func Detele(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	var msg struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
	}
	id := c.Param("id")
	topic, err := topicModel.GetTopicById(id)

	if err != nil {
		msg.Message = "此话题不存在"
		msg.Success = false
		c.JSON(http.StatusOK, msg)
		return
	}
	if topic.Author_id.Hex() != user.Id.Hex() {
		msg.Message = "无权限"
		msg.Success = false
		c.JSON(http.StatusForbidden, msg)
		return
	}

	topicModel.Delete(id)
	msg.Message = "话题删除成功"
	msg.Success = true
	c.JSON(http.StatusOK, msg)
}
func ShowEdit(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	var msg struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
	}
	id := c.Param("id")
	topic, err := topicModel.GetTopicById(id)

	if err != nil {
		msg.Message = "此话题不存在"
		msg.Success = false
		c.JSON(http.StatusNotFound, msg)
		return
	}
	if topic.Author_id.Hex() != user.Id.Hex() {
		msg.Message = "对不起，你不能编辑此话题"
		msg.Success = false
		c.JSON(http.StatusForbidden, msg)
		return
	}
	tabs := [3]map[string]string{{"value": "share", "text": "分享"}, {"value": "ask", "text": "问答"}, {"value": "job", "text": "招聘"}}
	c.HTML(http.StatusOK, "edit", gin.H{
		"user":     user,
		"action":   "edit",
		"topic_id": topic.Id.Hex(),
		"title":    topic.Title,
		"content":  topic.Content,
		"tab":      topic.Tab,
		"tabs":     tabs,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})

}
func Update(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	var msg struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
	}
	id := c.Param("id")
	topic, err := topicModel.GetTopicById(id)

	if err != nil {
		msg.Message = "此话题不存在"
		msg.Success = false
		c.JSON(http.StatusNotFound, msg)
		return
	}
	if topic.Author_id.Hex() != user.Id.Hex() {
		msg.Message = "对不起，你不能编辑此话题"
		msg.Success = false
		c.JSON(http.StatusForbidden, msg)
		return
	}
	tabs := [3]map[string]string{{"value": "share", "text": "分享"}, {"value": "ask", "text": "问答"}, {"value": "job", "text": "招聘"}}
	tab := c.Request.FormValue("tab")
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	// 验证
	editError := ""
	if title == "" {
		editError = "标题不能是空的。"
	}
	if len(title) < 5 || len(title) > 100 {
		editError = "标题字数太多或太少。"
	}
	if tab == "" {
		editError = "必须选择一个版块"
	}
	if content == "" {
		editError = "内容不可为空。"
	}
	if editError != "" {
		c.HTML(http.StatusOK, "edit", gin.H{
			"user":      user,
			"editError": editError,
			"action":    "edit",
			"topic_id":  topic.Id.Hex(),
			"title":     topic.Title,
			"content":   topic.Content,
			"tab":       topic.Tab,
			"tabs":      tabs,
			"config": gin.H{
				"description": "CNode：Node.js专业中文社区",
			},
		})
		return
	}
	topic.Title=title
	topic.Tab=tab
	topic.Content=content
	topicModel.Update(topic)
	url := "/topic/" + topic.Id.Hex()
	c.Redirect(301, url)

}
func Create(c *gin.Context) {
	session := sessions.Get(c)
	var name string
	user := models.User{}
	if nil != session.Get("loginname") {
		name = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(name)
	}
	id := user.Id.Hex()

	tab := c.Request.FormValue("tab")
	title := c.Request.FormValue("title")
	content := c.Request.FormValue("content")
	topic, _ := topicModel.NewAndSave(title, tab, id, content)
	url := "/topic/" + topic.Id.Hex()
	c.Redirect(301, url)
}

const (
	upload_path  string = "./public/upload/"
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

	id := user.Id.Hex()
	log.Println(id)
	//picName := c.Request.FormValue("name")
	file, _ := c.FormFile("file")

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
	c.SaveUploadedFile(file, upload_path+file.Filename)
	var msg struct {
		Success bool   `json:"success"`
		Url     string `json:"url"`
	}
	msg.Success = true
	msg.Url = upload_path2 + file.Filename
	c.JSON(http.StatusOK, msg)

}
