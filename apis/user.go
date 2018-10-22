package apis

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	. "github.com/dangyanglim/go_cnode/models"
	. "github.com/dangyanglim/go_cnode/utils"
	"github.com/gin-gonic/gin"
)

func IndexApi3(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}
func LoginAccess(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}
func GetUserCode(c *gin.Context) {
	phoneNumbers := c.Request.FormValue("mobile")
	accessKeyID := "LTAISKl55ZkE7jG2"
	accessSecret := "cHP1LfmXCA12sHWRlJwACXIyxE5Tr3"
	var numb = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var i = 0
	var token = ""
	rand.Seed(time.Now().Unix())
	for i = 0; i < 6; i++ {

		rnd := rand.Intn(10)
		token += strconv.Itoa(numb[rnd])
	}
	signName := "陈阳林"
	templateParam := `{"code":"` + token + `"}`
	templateCode := "SMS_117511815"
	now := time.Now()
	mm, _ := time.ParseDuration("10m")
	mm1 := now.Add(mm)
	tokenExptime := mm1.Format("2006-01-02 15:04:05")
	userCode := Code{Mobile: phoneNumbers, Token: token, TokenExptime: tokenExptime}
	userCodeTemp := Code{Mobile: phoneNumbers}
	code, err := userCodeTemp.GetUserCode()
	if err != nil {
		//log.Fatalln(err)

		log.Print(err.Error())
		ra, err2 := userCode.AddCode()
		if err != nil {
			//log.Fatalln(err)
			log.Println("add")
			log.Println(err.Error())
			err = err2
		}
		log.Println(ra)
	} else {
		log.Println("gengxin")
		ra, err := userCode.UpdateCode()
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
		}
		log.Println(ra, err)
	}
	log.Print(code, err)

	log.Println(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode)
	if err == nil {
		if err := SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode); err != nil {
			// 	log.Println("dysms.SendSms", err)
		}
	}
	retCode := ""
	msg := ""

	if err == nil {
		retCode = "0"
		msg = "发送验证码成功"
	} else {
		retCode = "7"
		msg = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"status":  "200",
		"code":    retCode,
		"result":  gin.H{},
	},
	)
}
func UserLoginByPwd(c *gin.Context) {
	mobile := c.Request.FormValue("mobile")
	pwd := c.Request.FormValue("pwd")
	log.Print(mobile)
	log.Print(pwd)
	log.Print("hh")
	ret := ""
	msg := "密码登陆成功"
	user := User{Mobile: mobile}
	userTemp, err := user.GetUser()
	log.Println(userTemp)
	if err != nil {
		ret = "2"
		msg = "没有该用户"
		log.Println(err.Error())
	} else {
		if userTemp.Pwd == pwd {
			ret = "0"
			msg = "密码登陆成功"
		} else {
			ret = "3"
			msg = "密码不对"
		}
	}
	if ret == "0" {
		var numb = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		var i = 0
		var token = ""
		rand.Seed(time.Now().Unix())
		for i = 0; i < 8; i++ {
			rnd := rand.Intn(10)
			token += strconv.Itoa(numb[rnd])
		}
		now := time.Now()
		mm, _ := time.ParseDuration("60m")
		mm1 := now.Add(mm)
		tokenExptime := mm1.Format("2006-01-02 15:04:05")
		userToken := UserToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
		ra, err := userToken.UpdateToken()
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
		}
		log.Println(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
			"status":  "200",
			"code":    ret,
			"result": gin.H{
				"token": token,
			},
		},
		)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"status":  "200",
		"code":    ret,
		"result": gin.H{
			"token": "kdkwi",
		},
	},
	)
}

type LoginByCode struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required"`
	CheckCode string `form:"checkCode" json:"checkCode" binding:"required"`
}

// type LoginByCode struct {
// 	User     string `form:"user" json:"user" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

func UserLoginByCode(c *gin.Context) {
	var form LoginByCode
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(err)

	mobile := c.Request.FormValue("mobile")
	checkCode := c.Request.FormValue("checkCode")
	log.Print(mobile)
	log.Print(checkCode)
	log.Print("hh")
	ret := ""
	msg := "验证码登陆成功"
	userCodeTemp := Code{Mobile: mobile}
	codeTemp, err := userCodeTemp.GetUserCode()
	if err != nil {
		log.Println(err.Error())
		ret = "2"
		msg = "没有该用户"
	} else {
		if codeTemp.Token == checkCode {
			ret = "0"
			msg = "验证码登陆成功"
		} else {
			ret = "1"
			msg = "验证码不对"
		}
	}
	userToken := UserToken{Mobile: mobile}
	var sendToken = ""
	userTokenTemp, err := userToken.GetUserToken()
	if err != nil {
		var numb = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		var i = 0
		var token = ""
		rand.Seed(time.Now().Unix())
		for i = 0; i < 8; i++ {

			rnd := rand.Intn(10)
			token += strconv.Itoa(numb[rnd])
		}
		now := time.Now()
		mm, _ := time.ParseDuration("60m")
		mm1 := now.Add(mm)
		tokenExptime := mm1.Format("2006-01-02 15:04:05")
		userToken = UserToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
		ra, err := userToken.AddToken()
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
		}
		log.Println(ra)
		sendToken = token
	} else {
		var numb = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		var i = 0
		var token = ""
		rand.Seed(time.Now().Unix())
		for i = 0; i < 8; i++ {
			rnd := rand.Intn(10)
			token += strconv.Itoa(numb[rnd])
		}
		now := time.Now()
		mm, _ := time.ParseDuration("60m")
		mm1 := now.Add(mm)
		tokenExptime := mm1.Format("2006-01-02 15:04:05")
		userToken = UserToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
		ra, err := userToken.UpdateToken()
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
		}
		log.Println(ra)
		sendToken = token
	}
	if sendToken == "" {
		sendToken = userTokenTemp.Token
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"status":  "200",
		"code":    ret,
		"result": gin.H{
			"token": sendToken,
		},
	},
	)
}

type UserRegisterType struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required"`
	CheckCode string `form:"checkCode" json:"checkCode" binding:"required"`
	Pwd       string `form:"pwd" json:"pwd" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var form UserRegisterType
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(err)
	mobile := c.Request.FormValue("mobile")
	checkCode := c.Request.FormValue("checkCode")
	pwd := c.Request.FormValue("pwd")
	log.Print(mobile)
	log.Print(checkCode)
	log.Print("hh")
	ret := ""
	msg := "验证码成功"
	userCodeTemp := Code{Mobile: mobile}
	codeTemp, err := userCodeTemp.GetUserCode()
	if err != nil {
		log.Println(err.Error())
		ret = "1"
		msg = "验证码不对"
	} else {
		if codeTemp.Token == checkCode {
			ret = "0"
			msg = "验证码成功"
		} else {
			ret = "1"
			msg = "验证码不对"
		}
	}
	if ret == "0" {
		userTemp := User{Mobile: mobile, Pwd: pwd}
		ra, err := userTemp.AddUser()
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
			msg = "创建用户失败"
			ret = "5"
		} else {
			ret = "0"
		}
		log.Println(ra)
		var numb = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		var i = 0
		var token = ""
		rand.Seed(time.Now().Unix())
		for i = 0; i < 8; i++ {

			rnd := rand.Intn(10)
			token += strconv.Itoa(numb[rnd])
		}
		now := time.Now()
		mm, _ := time.ParseDuration("60m")
		mm1 := now.Add(mm)
		tokenExptime := mm1.Format("2006-01-02 15:04:05")
		userToken := UserToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
		ra2, err2 := userToken.AddToken()
		if err2 != nil {
			//log.Fatalln(err)
			log.Println(err2.Error())
			msg = "创建用户Token失败"
			ret = "9"

			var numb = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			var i = 0
			var token = ""
			rand.Seed(time.Now().Unix())
			for i = 0; i < 8; i++ {
				rnd := rand.Intn(10)
				token += strconv.Itoa(numb[rnd])
			}
			now := time.Now()
			mm, _ := time.ParseDuration("60m")
			mm1 := now.Add(mm)
			tokenExptime := mm1.Format("2006-01-02 15:04:05")
			userToken = UserToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
			ra, err := userToken.UpdateToken()
			if err != nil {
				//log.Fatalln(err)
				log.Println(err.Error())
				ret = "6"
				msg = "更新用户Token失败"
			} else {
				ret = "0"
			}
			log.Println(ra)
		} else {
			ret = "0"
		}
		log.Println(ra2)
		if ret == "0" {
			msg = "注册成功"
			c.JSON(http.StatusOK, gin.H{
				"message": msg,
				"status":  "200",
				"code":    ret,
				"result": gin.H{
					"token": token,
				},
			},
			)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"status":  "200",
		"code":    ret,
		"result":  gin.H{},
	},
	)
}
func UserResetPwd(c *gin.Context) {
	var form UserRegisterType
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(err)
	mobile := c.Request.FormValue("mobile")
	checkCode := c.Request.FormValue("checkCode")
	pwd := c.Request.FormValue("pwd")
	log.Print(mobile)
	log.Print(checkCode)
	log.Print("hh")
	ret := ""
	msg := "验证码成功"
	userCodeTemp := Code{Mobile: mobile}
	codeTemp, err := userCodeTemp.GetUserCode()
	if err != nil {
		log.Println(err.Error())
		ret = "1"
		msg = "验证码不对"
	} else {
		if codeTemp.Token == checkCode {
			ret = "0"
			msg = "验证码成功"
		} else {
			ret = "1"
			msg = "验证码不对"
		}
	}
	if ret == "0" {
		userTemp := User{Mobile: mobile, Pwd: pwd}
		ra, err := userTemp.GetUser()
		var ra2 int64
		var err2 error
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
			msg = err.Error()
			ret = "5"
		} else {
			ra2, err2 = userTemp.UpdateUserPwd()
			if err2 != nil {
				ret = "8"
				msg = err.Error()
			}
		}
		log.Println(ra)

		log.Println(ra2)
		if err == nil && err2 == nil {
			msg = "重置成功"
			c.JSON(http.StatusOK, gin.H{
				"message": msg,
				"status":  "200",
				"code":    ret,
				"result":  gin.H{},
			},
			)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"status":  "200",
		"code":    ret,
		"result":  gin.H{},
	},
	)
}

type UserGetRoomType struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

func UserGetRoom(c *gin.Context) {
	var form UserGetRoomType
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(err)
	mobile := c.Request.FormValue("mobile")
	token := c.Request.FormValue("token")

	log.Print(mobile)
	log.Print(token)
	log.Print("hh")
	ret := ""
	msg := "验证码成功"
	userTokenTemp := UserToken{Mobile: mobile}
	tokenTemp, err := userTokenTemp.GetUserToken()
	if err != nil {
		log.Println(err.Error())
		ret = "9"
		msg = "token不存在"
	} else {
		//log.Print("token", tokenTemp.Token)
		if tokenTemp.Token == token {
			ret = "0"
			msg = "token成功"
		} else {
			ret = "1"
			msg = "token不对"
		}
	}
	if ret == "0" {
		renter := Renter{Renter: mobile, Status: "true"}
		var renters *[]map[string]string
		renters, err = renter.GetRenterByMobile()
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
			msg = "虚拟房源"
			ret = "0"
			c.JSON(http.StatusOK, gin.H{
				"message": msg,
				"status":  "200",
				"code":    ret,
				"result": gin.H{
					"roomname": "02",
					"floor":    "9",
					"building": "西区8",
					"garden":   "(虚拟)顺德碧桂园",
					"electric": gin.H{
						"now":       "2322.8",
						"nowMonth":  "30.2",
						"lastMonth": "39.8",
					},
					"water": gin.H{
						"now":       "88.7",
						"nowMonth":  "8.2",
						"lastMonth": "9.5",
					},
				},
			},
			)
			return
		} else if len(*renters) > 0 {
			log.Print(renters)
			log.Print((*renters)[0]["roomid"], (*renters)[0]["propertyid"])
			var propertyroom *[]map[string]string
			id, err := strconv.Atoi((*renters)[0]["roomid"])
			propertyroom, err = GetPropertyRoomById((*renters)[0]["propertyid"], id)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": err.Error(),
					"status":  "failed",
					"code":    "6",
					"result":  gin.H{},
				},
				)
			} else {
				log.Print(*propertyroom)
				c.JSON(http.StatusOK, gin.H{
					"message": "获取房间成功",
					"status":  "200",
					"code":    "0",
					"result": gin.H{
						"roomname":   (*propertyroom)[0]["room"],
						"floor":      (*propertyroom)[0]["floor"],
						"building":   (*propertyroom)[0]["building"],
						"garden":     (*propertyroom)[0]["garden"],
						"rentername": (*renters)[0]["name"],
						"electric": gin.H{
							"now":       "2322.8",
							"nowMonth":  "30.2",
							"lastMonth": "39.8",
						},
						"water": gin.H{
							"now":       "88.7",
							"nowMonth":  "8.2",
							"lastMonth": "9.5",
						},
					},
				},
				)
				return
			}
		}

	}
	msg = "虚拟房源"
	ret = "0"
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"status":  "200",
		"code":    ret,
		"result": gin.H{
			"roomname": "02",
			"floor":    "9",
			"building": "西区8",
			"garden":   "(虚拟)顺德碧桂园",
			"electric": gin.H{
				"now":       "2322.8",
				"nowMonth":  "30.2",
				"lastMonth": "39.8",
			},
			"water": gin.H{
				"now":       "88.7",
				"nowMonth":  "8.2",
				"lastMonth": "9.5",
			},
		},
	},
	)
}
