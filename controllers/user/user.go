package user

import (

	"net/http"
	//"regexp"
	"log"
	"go_cnode/mgoModels"
	//"github.com/dangyanglim/go_cnode/service/mail"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"time"
	"strconv"
	"math"

)

var userModel = new(models.UserModel)
var topicModel = new(models.TopicModel)
var replyModel = new(models.ReplyModel)


func Index(c *gin.Context) {
	session := sessions.Get(c)
	var loginname string

	user := models.User{}
	if nil != session.Get("loginname") {
		loginname = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(loginname)
	}
	name := c.Param("name")
	user, _ = userModel.GetUserByName(name)
	_,topics, _ := topicModel.GetAuthorTopics(user.Id.Hex(),5,0)
	var recent_topics []map[string]interface{}
	json.Unmarshal([]byte(topics), &recent_topics)

	for _, v := range recent_topics {
		timeString := v["topic"].(map[string]interface{})["Create_at"].(string)
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", timeString)
		v["topic"].(map[string]interface{})["Create_at"] = t.Format("2006-01-02 15:04:05")
		timeString = v["topic"].(map[string]interface{})["last_reply_at"].(string)
		t, _ = time.Parse("2006-01-02T15:04:05-07:00", timeString)
		v["topic"].(map[string]interface{})["last_reply_at"] = t.Format("2006-01-02 15:04:05")
		if v["reply"].(map[string]interface{})["Author_id"].(string) != "" {
			author, _ := userModel.GetUserById(v["reply"].(map[string]interface{})["Author_id"].(string))
			j, _ := json.Marshal(author)
			m := make(map[string]interface{})
			json.Unmarshal(j, &m)
			v["reply"].(map[string]interface{})["author"] = m
		}
	}
	topics2, _ := topicModel.GetReplyTopics(user.Id.Hex(),20,0,5)
	var recent_replies []map[string]interface{}
	json.Unmarshal([]byte(topics2), &recent_replies)
	for _, v := range recent_replies {
		timeString := v["topic"].(map[string]interface{})["Create_at"].(string)
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", timeString)
		v["topic"].(map[string]interface{})["Create_at"] = t.Format("2006-01-02 15:04:05")
		timeString = v["topic"].(map[string]interface{})["last_reply_at"].(string)
		t, _ = time.Parse("2006-01-02T15:04:05-07:00", timeString)
		v["topic"].(map[string]interface{})["last_reply_at"] = t.Format("2006-01-02 15:04:05")
		if v["reply"].(map[string]interface{})["Author_id"].(string) != "" {
			author, _ := userModel.GetUserById(v["reply"].(map[string]interface{})["Author_id"].(string))
			j, _ := json.Marshal(author)
			m := make(map[string]interface{})
			json.Unmarshal(j, &m)
			v["reply"].(map[string]interface{})["author"] = m
		}
	}		
	c.HTML(http.StatusOK, "userIndex", gin.H{
		"recent_topics":recent_topics,
		"recent_replies":recent_replies,
		"user":  user,		
	})
}
func Topics(c *gin.Context) {
	session := sessions.Get(c)
	var loginname string
	var pageSize = 20
	
	page := c.Request.FormValue("page")
	if page == "" {
		page = "1"
	}
	current_page, _ := strconv.Atoi(page)	
	user := models.User{}
	if nil != session.Get("loginname") {
		loginname = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(loginname)
	}
	name := c.Param("name")
	user, _ = userModel.GetUserByName(name)
	_,topics, _ := topicModel.GetAuthorTopics(user.Id.Hex(),pageSize,(current_page-1)*pageSize)
	var recent_topics []map[string]interface{}
	json.Unmarshal([]byte(topics), &recent_topics)

	for _, v := range recent_topics {
		timeString := v["topic"].(map[string]interface{})["Create_at"].(string)
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", timeString)
		v["topic"].(map[string]interface{})["Create_at"] = t.Format("2006-01-02 15:04:05")
		timeString = v["topic"].(map[string]interface{})["last_reply_at"].(string)
		t, _ = time.Parse("2006-01-02T15:04:05-07:00", timeString)
		v["topic"].(map[string]interface{})["last_reply_at"] = t.Format("2006-01-02 15:04:05")
		if v["reply"].(map[string]interface{})["Author_id"].(string) != "" {
			author, _ := userModel.GetUserById(v["reply"].(map[string]interface{})["Author_id"].(string))
			j, _ := json.Marshal(author)
			m := make(map[string]interface{})
			json.Unmarshal(j, &m)
			v["reply"].(map[string]interface{})["author"] = m
		}
	}
	var page_start int
	pages, _:= topicModel.GetTopicByAuthorQueryCount(user.Id,"all", false)
	pages = int(math.Floor(float64(pages)/float64(pageSize))) + 1
	base_url := "?page="
	var page_end int

	log.Println(current_page)	
	if (current_page - 2) > 0 {
		page_start = current_page - 2
	} else {
		page_start = 1
	}
	if (page_start + 4) > pages {
		page_end = pages
	} else {
		page_end = page_start + 4
	}
	pagesArray := []int{}
	var i int
	for i = 1; i < pages+1; i++ {
		pagesArray = append(pagesArray, i)
	}	
	//log.Println(recent_topics)
	c.HTML(http.StatusOK, "userTopics", gin.H{
		"user":  user,
		"topicss":recent_topics,
		"pages":           pages,
		"page_start":      page_start,
		"page_end":        page_end,
		"pagesArray":      pagesArray,
		"base_url":        base_url,
		"current_page":    current_page,		
	})	

}
func Replies(c *gin.Context) {
	session := sessions.Get(c)
	var loginname string
	var pageSize = 20
	
	page := c.Request.FormValue("page")
	if page == "" {
		page = "1"
	}
	current_page, _ := strconv.Atoi(page)	
	user := models.User{}
	if nil != session.Get("loginname") {
		loginname = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(loginname)
	}
	name := c.Param("name")
	user, _ = userModel.GetUserByName(name)
	topics2, _ := topicModel.GetReplyTopics(user.Id.Hex(),pageSize,(current_page-1)*pageSize,pageSize)
	var recent_replies []map[string]interface{}
	json.Unmarshal([]byte(topics2), &recent_replies)
	for _, v := range recent_replies {
		timeString := v["topic"].(map[string]interface{})["Create_at"].(string)
		t, _ := time.Parse("2006-01-02T15:04:05-07:00", timeString)
		v["topic"].(map[string]interface{})["Create_at"] = t.Format("2006-01-02 15:04:05")
		timeString = v["topic"].(map[string]interface{})["last_reply_at"].(string)
		t, _ = time.Parse("2006-01-02T15:04:05-07:00", timeString)
		v["topic"].(map[string]interface{})["last_reply_at"] = t.Format("2006-01-02 15:04:05")
		if v["reply"].(map[string]interface{})["Author_id"].(string) != "" {
			author, _ := userModel.GetUserById(v["reply"].(map[string]interface{})["Author_id"].(string))
			j, _ := json.Marshal(author)
			m := make(map[string]interface{})
			json.Unmarshal(j, &m)
			v["reply"].(map[string]interface{})["author"] = m
		}
	}	

	var page_start int
	pages, _:= replyModel.GetReplyByAuthorQueryCount(user.Id)
	pages = int(math.Floor(float64(pages)/float64(pageSize))) + 1
	base_url := "?page="
	var page_end int

	log.Println(current_page)	
	if (current_page - 2) > 0 {
		page_start = current_page - 2
	} else {
		page_start = 1
	}
	if (page_start + 4) > pages {
		page_end = pages
	} else {
		page_end = page_start + 4
	}
	pagesArray := []int{}
	var i int
	for i = 1; i < pages+1; i++ {
		pagesArray = append(pagesArray, i)
	}	
	//log.Println(recent_topics)
	c.HTML(http.StatusOK, "userReplies", gin.H{
		"user":  user,
		"topicss":recent_replies,
		"pages":           pages,
		"page_start":      page_start,
		"page_end":        page_end,
		"pagesArray":      pagesArray,
		"base_url":        base_url,
		"current_page":    current_page,		
	})

}
func Top100(c *gin.Context) {
	session := sessions.Get(c)
	var loginname string

	user := models.User{}
	if nil != session.Get("loginname") {
		loginname = session.Get("loginname").(string)
		user, _ = userModel.GetUserByName(loginname)
	}
	tops, _ := userModel.GetUserTops(100)
	c.HTML(http.StatusOK, "userTop100", gin.H{
		"user":  user,
		"users":tops,		
	})	
}
