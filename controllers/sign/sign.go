package sign

import (
	"encoding/json"
	"go_cnode/mgoModels"
	"go_cnode/service/mail"
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
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
	pass := c.Request.FormValue("pass")

	user, _ := userModel.GetUserByName(name)
	err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(pass))
	log.Println(err)
	if err != nil || !user.Active {

		c.Redirect(301, "/")
		return
	}
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
	userModel.NewAndSave(loginname, loginname, email, pass, "aa", false)
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

func GithubSignup(c *gin.Context) {
	client_id := "bafc506847f325223094"
	client_secret := "172d5424cc25be8c8ac8095b40d79fea859588ea"
	AuthURL := "https://github.com/login/oauth/authorize?"

	url := AuthURL + "client_id=" + client_id + "&client_secret=" + client_secret
	log.Println(url)
	c.Redirect(http.StatusMovedPermanently, url)

}

type GithubUser struct {
	Login      string `json:"login"`
	Id         int    `json:"id"`
	Avatar_url string `json:"avatar_url"`
	Email      string `json:"email"`
}
type Token struct {
	Access_token string `json:"access_token"`
}

func GithubCallBack(c *gin.Context) {
	TokenURL := "https://github.com/login/oauth/access_token?"
	UserURL := "https://api.github.com/user?"
	client_id := "bafc506847f325223094"
	client_secret := "172d5424cc25be8c8ac8095b40d79fea859588ea"
	code := c.Request.FormValue("code")
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	url := TokenURL + "client_id=" + client_id + "&client_secret=" + client_secret + "&code=" + code
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	res, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// handle error
	}
	var token Token
	json.Unmarshal(body, &token)
	log.Println(string(body))
	log.Println(token)

	url = UserURL + "access_token=" + token.Access_token
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	res, err = httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		// handle error
	}
	var githubUser GithubUser
	json.Unmarshal(body, &githubUser)
	log.Println(string(body))
	log.Println(githubUser)
	user, er := userModel.GetUserByGithubId(githubUser.Id)
	//user, er := userModel.GetUserByName("admin")
	log.Println(er)
	log.Println(user)
	if er == nil {
		session := sessions.Get(c)
		session.Set("loginname", user.Loginname)
		session.Set("accessToken", user.AccessToken)
		session.Save()
		c.Redirect(301, "/")
		//return
	} else {
		user2, er2 := userModel.GithubNewAndSave(githubUser.Login, githubUser.Login, githubUser.Email, githubUser.Avatar_url, true, githubUser.Id)
		session := sessions.Get(c)
		session.Set("loginname", user2.Loginname)
		session.Set("accessToken", user2.AccessToken)
		session.Save()
		log.Println(user2)
		log.Println(er2)
		c.Redirect(301, "/")
	}
	//c.JSON(200, githubUser)
}
func ActiveAccount(c *gin.Context) {
	loginname := c.Request.FormValue("name")

	user := models.User{}
	var err error

	user, err = userModel.GetUserByName(loginname)
	if err != nil {
		c.HTML(http.StatusOK, "notify", gin.H{
			"error": "用户不存在",
			"config": gin.H{
				"description": "CNode：Node.js专业中文社区",
			},
		})
		return
	}
	userModel.ActiveUserByName(loginname)
	log.Println(user)
	c.HTML(http.StatusOK, "notify", gin.H{
		"success": "帐号已被激活，请登录",
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
