package apis

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	. "github.com/dangyanglim/go_cnode/models"
	. "github.com/dangyanglim/go_cnode/utils"
	"github.com/gin-gonic/gin"
)

func sendUserNotice(mobile string, addr string, timeBegin string, timeEnd string, pwd string) (err error) {
	phoneNumbers := mobile
	accessKeyID := "LTAISKl55ZkE7jG2"
	accessSecret := "cHP1LfmXCA12sHWRlJwACXIyxE5Tr3"

	signName := "陈阳林"
	templateParam := `{"addr":"` + addr + `"` + `,"begin":"` + timeBegin + `"` + `,"end":"` + timeEnd + `"` + `,"pwd":"` + pwd + `"}`
	templateCode := "SMS_130913665"

	//templateParam := `{"addr":"12345","date":"hhhh","pwd":"hh"}`
	log.Println(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode)
	if err = SendSms(accessKeyID, accessSecret, phoneNumbers, signName, templateParam, templateCode); err != nil {
		log.Println("dysms.SendSms", err)
	}
	return err

}

func IndexApi4(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func GetPropertyCode(c *gin.Context) {
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
	propertyCode := PropertyCheckCode{Mobile: phoneNumbers, CheckCode: token, TokenExptime: tokenExptime}
	propertyCodeTemp := PropertyCheckCode{Mobile: phoneNumbers}
	code, err := propertyCodeTemp.GetPropertyCheckCode()
	if err != nil {
		//log.Fatalln(err)

		log.Print(err.Error())
		ra, err2 := propertyCode.AddPropertyCheckCode()
		if err != nil {
			//log.Fatalln(err)
			log.Println("add")
			log.Println(err.Error())
			err = err2
		}
		log.Println(ra)
	} else {
		log.Println("gengxin")
		ra, err := propertyCode.UpdatePropertyCode()
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
func PropertyLoginByPwd(c *gin.Context) {
	mobile := c.Request.FormValue("mobile")
	pwd := c.Request.FormValue("pwd")
	log.Print(mobile)
	log.Print(pwd)
	log.Print("hh")
	ret := ""
	msg := "密码登陆成功"
	property := Property{Mobile: mobile}
	propertyTemp, err := property.GetProperty()
	log.Println(propertyTemp)
	log.Print("loginbypwd")
	if err != nil {
		ret = "2"
		msg = "没有该用户"
		log.Println(err.Error())
	} else {
		if propertyTemp.Pwd == pwd {
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
		propertyToken := PropertyToken{Mobile: mobile}
		propertyTokenTemp, err := propertyToken.GetPropertyToken()
		if err != nil {
			ret = "2"
			msg = "没有该用户"
			log.Println(err.Error())
			log.Print("no token")
			log.Print(propertyTokenTemp)
			propertyToken := PropertyToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
			ra, err2 := propertyToken.AddPropertyToken()
			log.Print(ra)
			log.Print(err2)

		} else {
			propertyToken := PropertyToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
			ra, err := propertyToken.UpdatePropertyToken()
			if err != nil {
				//log.Fatalln(err)
				log.Println(err.Error())
			}
			log.Println(ra)
		}

		if ret == "0" {
			c.JSON(http.StatusOK, gin.H{
				"message":          msg,
				"status":           "ok",
				"code":             ret,
				"type":             "account",
				"currentAuthority": "admin",
				"result": gin.H{
					"token":  token,
					"mobile": mobile,
					"name":   propertyTemp.Name,
				},
			},
			)
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"status":  "failed",
		"code":    ret,
		"result": gin.H{
			"token": "kdkwi",
		},
	},
	)
}
func GetProperty(c *gin.Context) {
	mobile := c.Request.FormValue("mobile")
	token := c.Request.FormValue("token")
	log.Print(mobile)
	log.Print(token)
	log.Print("hh")
	ret := ""
	msg := "密码登陆成功"
	propertyToken := PropertyToken{Mobile: mobile}
	propertyTokenTemp, err := propertyToken.GetPropertyToken()
	log.Println(propertyTokenTemp)
	if err != nil {
		ret = "2"
		msg = "没有该用户"
		log.Println(err.Error())
		log.Print("no token")
	} else {
		if propertyTokenTemp.Token == token {
			ret = "0"
			msg = "token验证成功"
		} else {
			ret = "3"
			msg = "token不对"
		}
	}
	if ret == "0" {

		property := Property{Mobile: mobile}
		ra, err := property.GetProperty()
		log.Println(ra)
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
			ret = "2"
			msg = "没有该用户"
		} else {
			var name string
			if ra.Name != "" {
				name = ra.Name
			} else {
				name = ra.Mobile
			}
			c.JSON(http.StatusOK, gin.H{
				"message":          msg,
				"status":           "ok",
				"code":             ret,
				"type":             "account",
				"currentAuthority": "admin",
				"name":             name,
				"result": gin.H{
					"token": token,
				},
			},
			)
		}

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"status":  "failed",
		"code":    ret,
		"result": gin.H{
			"token": "kdkwi",
		},
	},
	)
}

type PropertyCode struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required"`
	CheckCode string `form:"checkCode" json:"checkCode" binding:"required"`
}

// type LoginByCode struct {
// 	User     string `form:"user" json:"user" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

func PropertyLoginByCode(c *gin.Context) {
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
	propertyCodeTemp := PropertyCheckCode{Mobile: mobile}
	codeTemp, err := propertyCodeTemp.GetPropertyCheckCode()
	if err != nil {
		log.Println(err.Error())
		ret = "2"
		msg = "没有该用户"
	} else {
		if codeTemp.CheckCode == checkCode {
			ret = "0"
			msg = "验证码登陆成功"
		} else {
			ret = "1"
			msg = "验证码不对"
		}
	}
	propertyToken := PropertyToken{Mobile: mobile}
	var sendToken = ""
	propertyTokenTemp, err := propertyToken.GetPropertyToken()
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
		propertyToken = PropertyToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
		ra, err := propertyToken.AddPropertyToken()
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
		propertyToken = PropertyToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
		ra, err := propertyToken.UpdatePropertyToken()
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
		}
		log.Println(ra)
		sendToken = token
	}
	if sendToken == "" {
		sendToken = propertyTokenTemp.Token
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

type PropertyRegisterType struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required"`
	CheckCode string `form:"checkCode" json:"checkCode" binding:"required"`
	Pwd       string `form:"pwd" json:"pwd" binding:"required"`
}

func PropertyRegister(c *gin.Context) {
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
	propertyCodeTemp := PropertyCheckCode{Mobile: mobile}
	codeTemp, err := propertyCodeTemp.GetPropertyCheckCode()
	if err != nil {
		log.Println(err.Error())
		ret = "1"
		msg = "验证码不对"
	} else {
		if codeTemp.CheckCode == checkCode {
			ret = "0"
			msg = "验证码成功"
		} else {
			ret = "1"
			msg = "验证码不对"
		}
	}
	if ret == "0" {
		propertyTemp := Property{Mobile: mobile, Pwd: pwd}
		ra, err := propertyTemp.AddProperty()
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
			msg = "创建用户失败"
			ret = "5"
			c.JSON(http.StatusOK, gin.H{
				"message": msg,
				"status":  "failed",
				"code":    ret,
				"result":  gin.H{},
			},
			)
			return
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
		propertyToken := PropertyToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
		ra2, err2 := propertyToken.AddPropertyToken()
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
			propertyToken = PropertyToken{Mobile: mobile, Token: token, TokenExptime: tokenExptime}
			ra, err := propertyToken.UpdatePropertyToken()
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
				"status":  "ok",
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
		"status":  "failed",
		"code":    ret,
		"result":  gin.H{},
	},
	)
}
func PropertyResetPwd(c *gin.Context) {
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
	propertyCodeTemp := PropertyCheckCode{Mobile: mobile}
	codeTemp, err := propertyCodeTemp.GetPropertyCheckCode()
	if err != nil {
		log.Println(err.Error())
		ret = "1"
		msg = "验证码不对"
	} else {
		if codeTemp.CheckCode == checkCode {
			ret = "0"
			msg = "验证码成功"
		} else {
			ret = "1"
			msg = "验证码不对"
		}
	}
	if ret == "0" {
		propertyTemp := Property{Mobile: mobile, Pwd: pwd}
		ra, err := propertyTemp.GetProperty()
		var ra2 int64
		var err2 error
		if err != nil {
			//log.Fatalln(err)
			log.Println(err.Error())
			msg = err.Error()
			ret = "5"
		} else {
			ra2, err2 = propertyTemp.UpdatePropertyPwd()
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
				"status":  "ok",
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

type PropertyGetRoomType struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
}

func checkToken(c *gin.Context) (Property, error) {
	var property Property
	mobile := c.Request.FormValue("mobile")
	token := c.Request.FormValue("token")
	log.Print(mobile)
	log.Print(token)
	propertyToken := PropertyToken{Mobile: mobile}
	propertyTokenTemp, err := propertyToken.GetPropertyToken()
	log.Println(propertyTokenTemp)
	if err != nil {
		return property, err
	} else {
		if propertyTokenTemp.Token == token {
			property = Property{Mobile: mobile}
			property, err := property.GetProperty()
			if err != nil {
				return property, errors.New("不存在")
			} else {
				return property, nil
			}

		} else {
			return property, errors.New("token不对")
		}
	}
}

type PropertyRoomType struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Token    string `form:"token" json:"token" binding:"required"`
	Floor    string `json:"floor" form:"floor" binding:"required"`
	Room     string `json:"room" form:"room" binding:"required"`
	Building string `json:"building" form:"building" binding:"required"`
	Garden   string `json:"garden" form:"garden" binding:"required"`
}

func AddPropertyRoomApi(c *gin.Context) {
	var form PropertyRoomType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(form.Room)
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	propertyRoom := PropertyRoom{Propertyid: s, Room: form.Room, Floor: form.Floor, Garden: form.Garden, Building: form.Building}
	err = propertyRoom.AddPropertyRoom()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "创建房间成功",
			"status":  "200",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}

type GetPropertyRoomType struct {
	Mobile      string `form:"mobile" json:"mobile" binding:"required"`
	Token       string `form:"token" json:"token" binding:"required"`
	CurrentPage int    `form:"currentPage" json:"currentPage" binding:"required"`
	PageSize    int    `form:"pageSize" json:"pageSize" binding:"required"`
}

func GetPropertyRoomApi(c *gin.Context) {
	var form GetPropertyRoomType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	var propertyroom *[]map[string]string
	var start = (form.CurrentPage - 1) * form.PageSize
	var end = form.CurrentPage * form.PageSize
	propertyroom, err = GetPropertyRoom(s, start, end)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(*propertyroom)
		//pageSize, err := strconv.Atoi(form.PageSize)
		//currentPage, err := strconv.Atoi(form.CurrentPage)
		count, err := GetRoomCount(s)
		total, err := strconv.Atoi(count)
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "获取房源成功",
			"status":  "200",
			"code":    "0",
			"result": gin.H{
				"room": gin.H{
					"list": *propertyroom,
					"pagination": gin.H{
						"total":    total,
						"pagesize": form.PageSize,
						"current":  form.CurrentPage,
					},
				},
			},
		},
		)
	}
}

type GetPropertyRoomByIdType struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
	Id     int    `json:"id" form:"id" binding:"required"`
}

func GetPropertyRoomByIdApi(c *gin.Context) {
	var form GetPropertyRoomByIdType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	var propertyroom *[]map[string]string
	propertyroom, err = GetPropertyRoomById(s, form.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(*propertyroom)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  "200",
				"code":    "4",
				"result":  gin.H{},
			},
			)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "获取房源成功",
			"status":  "200",
			"code":    "0",
			"result": gin.H{
				"room": *propertyroom,
			},
		},
		)
	}
}
func DelRoomApi(c *gin.Context) {
	var form GetPropertyRoomByIdType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	room := Room{Propertyid: s, Id: form.Id}
	ra, err := room.DelRoom()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "删除成功",
			"status":  "200",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}

}

type PropertyRenterType struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required"`
	Token     string `form:"token" json:"token" binding:"required"`
	Roomname  string `json:"roomname" form:"roomname" binding:"required"`
	Roomid    string `json:"roomid" form:"roomid" binding:"required"`
	Renter    string `json:"renter" form:"renter" binding:"required"`
	Name      string `json:"name" form:"name" binding:"required"`
	TimeBegin string `json:"timeBegin" form:"timeBegin" binding:"required"`
	TimeEnd   string `json:"timeEnd" form:"timeEnd" binding:"required"`
}

func AddPropertyRenterApi(c *gin.Context) {
	var form PropertyRenterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Print(form.Room)
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	renter := Renter{Propertyid: s, Renter: form.Renter, Name: form.Name, TimeBegin: form.TimeBegin, TimeEnd: form.TimeEnd, Roomid: form.Roomid, Status: "true"}
	id, err := renter.AddRenter()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(id)
		var numb = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		var i = 0
		var token = ""
		rand.Seed(time.Now().Unix())
		for i = 0; i < 8; i++ {

			rnd := rand.Intn(10)
			token += strconv.Itoa(numb[rnd])
		}
		err = sendUserNotice(form.Renter, form.Roomname, form.TimeBegin, form.TimeEnd, token)
		if err != nil {
			log.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "增加住户成功",
			"status":  "200",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}

type GetPropertyRenterType struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Token  string `form:"token" json:"token" binding:"required"`
	Roomid string `json:"roomid" form:"roomid" binding:"required"`
}

func GetPropertyRentersApi(c *gin.Context) {
	var form GetPropertyRenterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	var renters *[]map[string]string
	renter := Renter{Propertyid: s, Roomid: form.Roomid, Status: "true"}
	renters, err = renter.GetRenterByPropertyidByRoomid()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
				"status":  "200",
				"code":    "4",
				"result":  gin.H{},
			},
			)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "获取住户成功",
			"status":  "200",
			"code":    "0",
			"result": gin.H{
				"renters": gin.H{
					"list": *renters,
					"pagination": gin.H{
						"total":    5,
						"pagesize": 10,
						"current":  1,
					},
				},
			},
		},
		)
	}
}

type DelPropertyRenterType struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Token    string `form:"token" json:"token" binding:"required"`
	Renterid string `json:"renterid" form:"renterid" binding:"required"`
	Roomid   string `json:"roomid" form:"roomid" binding:"required"`
}

func DelPropertyRenterApi(c *gin.Context) {
	var form DelPropertyRenterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	// var renters *[]map[string]string
	id, err := strconv.Atoi(form.Renterid)
	renter := Renter{Propertyid: s, Id: id, Roomid: form.Roomid, Status: "false"}
	ra, err := renter.SetRenterStatus()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "5",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "修改住户成功",
			"status":  "200",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}

func GetPropertRoomCountApi(c *gin.Context) {
	type GetPropertyRoom struct {
		Mobile string `form:"mobile" json:"mobile" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
	}
	var form GetPropertyRoom
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	count, err := GetRoomCount(s)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "获取房间总数成功",
			"status":  "200",
			"code":    "0",
			"result": gin.H{
				"count": count,
			},
		},
		)
	}
}

func AddAmmeterApi(c *gin.Context) {
	type AddAmmeterType struct {
		Mobile       string `form:"mobile" json:"mobile" binding:"required"`
		Token        string `form:"token" json:"token" binding:"required"`
		Spaceid      string `json:"spaceid" form:"spaceid" binding:"required"`
		Spacetype    string `json:"spacetype" form:"spacetype" binding:"required"`
		Ammeter_addr string `json:"ammeter_addr" form:"ammeter_addr" binding:"required"`
		GateWay      string `json:"gateway" form:"gateway" binding:"required"`
	}
	var form AddAmmeterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	gateway := Gateway{Propertyid: s, GateWay: form.GateWay}
	gateway, err = gateway.GetGateway()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "采集器不存在",
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
		return
	} else {
		log.Print(gateway)
	}
	ammeter := Ammeter{Spaceid: form.Spaceid, Spacetype: form.Spacetype, Propertyid: s}
	count, err := ammeter.GetAmmeterCount()
	if err == nil && count > 0 {
		log.Print(count)
		c.JSON(http.StatusOK, gin.H{
			"message": "只能关联一个电表",
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
		return
	}
	ammeter = Ammeter{GateWay: form.GateWay, Spaceid: form.Spaceid, Spacetype: form.Spacetype, Ammeter_addr: form.Ammeter_addr, Propertyid: s, Voltage: "0", Current: "0", Energy: "0"}
	ra, err := ammeter.AddAmmeter()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "新增电表成功",
			"status":  "ok",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}
func GetAmmetersApi(c *gin.Context) {
	type AddAmmeterType struct {
		Mobile    string `form:"mobile" json:"mobile" binding:"required"`
		Token     string `form:"token" json:"token" binding:"required"`
		Spaceid   string `json:"spaceid" form:"spaceid" binding:"required"`
		Spacetype string `json:"spacetype" form:"spacetype" binding:"required"`
	}
	var form AddAmmeterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	ammeter := Ammeter{Spaceid: form.Spaceid, Spacetype: form.Spacetype, Propertyid: s}
	ammeters, err := ammeter.GetAmmeters()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "获取电表成功",
			"status":  "ok",
			"code":    "0",
			"result": gin.H{
				"ammeters": gin.H{
					"list": *ammeters,
					"pagination": gin.H{
						"total":    5,
						"pagesize": 10,
						"current":  1,
					},
				},
			},
		},
		)
	}
}
func DelAmmeterApi(c *gin.Context) {
	type DelAmmeterType struct {
		Mobile string `form:"mobile" json:"mobile" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
		Id     string `json:"id" form:"id" binding:"required"`
	}
	var form DelAmmeterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	// s := strconv.Itoa(property.Id)
	id, err := strconv.Atoi(form.Id)
	ammeter := Ammeter{Id: id}
	ra, err := ammeter.DelAmmeter()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "删除电表成功",
			"status":  "ok",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}
func AddLockApi(c *gin.Context) {
	type AddLockType struct {
		Mobile    string `form:"mobile" json:"mobile" binding:"required"`
		Token     string `form:"token" json:"token" binding:"required"`
		Spaceid   string `json:"spaceid" form:"spaceid" binding:"required"`
		Spacetype string `json:"spacetype" form:"spacetype" binding:"required"`
		Lock_addr string `json:"lock_addr" form:"lock_addr" binding:"required"`
		GateWay   string `json:"gateway" form:"gateway" binding:"required"`
	}
	var form AddLockType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	lock := Lock{GateWay: form.GateWay, Spaceid: form.Spaceid, Spacetype: form.Spacetype, Lock_addr: form.Lock_addr, Propertyid: s}
	ra, err := lock.AddLock()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "新增门锁成功",
			"status":  "ok",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}
func GetLocksApi(c *gin.Context) {
	type GetLocksType struct {
		Mobile    string `form:"mobile" json:"mobile" binding:"required"`
		Token     string `form:"token" json:"token" binding:"required"`
		Spaceid   string `json:"spaceid" form:"spaceid" binding:"required"`
		Spacetype string `json:"spacetype" form:"spacetype" binding:"required"`
	}
	var form GetLocksType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	lock := Lock{Spaceid: form.Spaceid, Spacetype: form.Spacetype, Propertyid: s}
	locks, err := lock.GetLocks()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "获取门锁成功",
			"status":  "ok",
			"code":    "0",
			"result": gin.H{
				"locks": gin.H{
					"list": *locks,
					"pagination": gin.H{
						"total":    5,
						"pagesize": 10,
						"current":  1,
					},
				},
			},
		},
		)
	}
}
func DelLockApi(c *gin.Context) {
	type DelLockType struct {
		Mobile string `form:"mobile" json:"mobile" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
		Id     string `json:"id" form:"id" binding:"required"`
	}
	var form DelLockType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	// s := strconv.Itoa(property.Id)
	id, err := strconv.Atoi(form.Id)
	lock := Lock{Id: id}
	ra, err := lock.DelLock()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "删除门锁成功",
			"status":  "ok",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}
func AddWaterApi(c *gin.Context) {
	type AddWaterType struct {
		Mobile     string `form:"mobile" json:"mobile" binding:"required"`
		Token      string `form:"token" json:"token" binding:"required"`
		Spaceid    string `json:"spaceid" form:"spaceid" binding:"required"`
		Spacetype  string `json:"spacetype" form:"spacetype" binding:"required"`
		Water_addr string `json:"water_addr" form:"water_addr" binding:"required"`
		GateWay    string `json:"gateway" form:"gateway" binding:"required"`
	}
	var form AddWaterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	water := Water{GateWay: form.GateWay, Spaceid: form.Spaceid, Spacetype: form.Spacetype, Water_addr: form.Water_addr, Propertyid: s}
	ra, err := water.AddWater()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "新增水表成功",
			"status":  "ok",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}
func GetWatersApi(c *gin.Context) {
	type GetWatersType struct {
		Mobile    string `form:"mobile" json:"mobile" binding:"required"`
		Token     string `form:"token" json:"token" binding:"required"`
		Spaceid   string `json:"spaceid" form:"spaceid" binding:"required"`
		Spacetype string `json:"spacetype" form:"spacetype" binding:"required"`
	}
	var form GetWatersType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	lock := Water{Spaceid: form.Spaceid, Spacetype: form.Spacetype, Propertyid: s}
	waters, err := lock.GetWaters()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "获取水表成功",
			"status":  "ok",
			"code":    "0",
			"result": gin.H{
				"waters": gin.H{
					"list": *waters,
					"pagination": gin.H{
						"total":    5,
						"pagesize": 10,
						"current":  1,
					},
				},
			},
		},
		)
	}
}
func DelWaterApi(c *gin.Context) {
	type DelWaterType struct {
		Mobile string `form:"mobile" json:"mobile" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
		Id     string `json:"id" form:"id" binding:"required"`
	}
	var form DelWaterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	// s := strconv.Itoa(property.Id)
	id, err := strconv.Atoi(form.Id)
	water := Water{Id: id}
	ra, err := water.DelWater()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "删除水表成功",
			"status":  "ok",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}
func AddGatewayApi(c *gin.Context) {
	type AddGatewayType struct {
		Mobile    string `form:"mobile" json:"mobile" binding:"required"`
		Token     string `form:"token" json:"token" binding:"required"`
		Spaceid   string `json:"spaceid" form:"spaceid" binding:"required"`
		Spacetype string `json:"spacetype" form:"spacetype" binding:"required"`
		GateWay   string `json:"gateway" form:"gateway" binding:"required"`
		Name      string `json:"name" form:"name" binding:"required"`
	}
	var form AddGatewayType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	gateway := Gateway{GateWay: form.GateWay, Spaceid: form.Spaceid, Spacetype: form.Spacetype, Propertyid: s, Name: form.Name}
	ra, err := gateway.AddGateway()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "新增网关成功",
			"status":  "ok",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}
func GetGatewaysApi(c *gin.Context) {
	type GetGatewaysType struct {
		Mobile    string `form:"mobile" json:"mobile" binding:"required"`
		Token     string `form:"token" json:"token" binding:"required"`
		Spaceid   string `json:"spaceid" form:"spaceid" binding:"required"`
		Spacetype string `json:"spacetype" form:"spacetype" binding:"required"`
	}
	var form GetGatewaysType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	s := strconv.Itoa(property.Id)
	gateway := Gateway{Spaceid: form.Spaceid, Spacetype: form.Spacetype, Propertyid: s}
	gateways, err := gateway.GetGateways()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "获取网关成功",
			"status":  "ok",
			"code":    "0",
			"result": gin.H{
				"gateways": gin.H{
					"list": *gateways,
					"pagination": gin.H{
						"total":    5,
						"pagesize": 10,
						"current":  1,
					},
				},
			},
		},
		)
	}
}
func DelGatewayApi(c *gin.Context) {
	type DelWaterType struct {
		Mobile string `form:"mobile" json:"mobile" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
		Id     string `json:"id" form:"id" binding:"required"`
	}
	var form DelWaterType
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	// s := strconv.Itoa(property.Id)
	id, err := strconv.Atoi(form.Id)
	gateway := Gateway{Id: id}
	ra, err := gateway.DelGateway()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "删除网关成功",
			"status":  "ok",
			"code":    "0",
			"result":  gin.H{},
		},
		)
	}
}
func UpdatePropertyNameApi(c *gin.Context) {
	type Request struct {
		Mobile string `form:"mobile" json:"mobile" binding:"required"`
		Token  string `form:"token" json:"token" binding:"required"`
		Name   string `json:"name" form:"name" binding:"required"`
	}
	var form Request
	var property Property
	var err error
	if err = c.ShouldBind(&form); err == nil {
		log.Print("hh2")

	} else {
		log.Print("hh3")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if property, err = checkToken(c); err != nil {
		log.Println(property)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "200",
			"code":    "3",
			"result":  gin.H{},
		},
		)
		return
	}
	// s := strconv.Itoa(property.Id)
	// id, err := strconv.Atoi(form.Id)
	property = Property{Mobile: form.Mobile, Name: form.Name}
	ra, err := property.UpdatePropertyName()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"status":  "failed",
			"code":    "4",
			"result":  gin.H{},
		},
		)
	} else {
		log.Print(ra)
		c.JSON(http.StatusOK, gin.H{
			"message": "更新名称成功",
			"status":  "ok",
			"code":    "0",
			"result": gin.H{
				"name": form.Name,
			},
		},
		)
	}
}
