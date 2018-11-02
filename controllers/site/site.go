package site

import (
	"net/http"
  "log"
	//. "github.com/dangyanglim/go_cnode/models"
	"html/template"
	"github.com/gin-gonic/gin"
  "gopkg.in/russross/blackfriday.v2"
  //db "github.com/dangyanglim/go_cnode/database"
  "github.com/tommy351/gin-sessions"
  "github.com/dangyanglim/go_cnode/mgoModels"
  "strconv"
  "github.com/dangyanglim/go_cnode/service/cache"
  "encoding/json"

)

var api =[]byte(`

以下 api 路径均以 **https://cnodejs.org/api/v1** 为前缀

### 主题

#### get /topics 主题首页

接收 get 参数

* page Number 页数
* tab String 主题分类。目前有 ask share job good
* limit Number 每一页的主题数量
* mdrender String 当为 false 时，不渲染。默认为 true，渲染出现的所有 markdown 格式文本。

示例：[/api/v1/topics](/api/v1/topics)

#### get /topic/:id 主题详情

接收 get 参数

* mdrender String 当为 alse 时，不渲染。默认为 true，渲染出现的所有 markdown 格式文本。
* accesstoken String 当需要知道一个主题是否被特定用户收藏以及对应评论是否被特定用户点赞时，才需要带此参数。会影响返回值中的 is_collect 以及 replies 列表中的 is_uped 值。

示例：[/api/v1/topic/5433d5e4e737cbe96dcef312](/api/v1/topic/5433d5e4e737cbe96dcef312)

#### post /topics 新建主题

接收 post 参数

* accesstoken String 用户的 accessToken
* title String 标题
* tab String 目前有 ask share job dev。开发新客户端的同学，请务必将你们的测试帖发在 dev 专区，以免污染日常的版面，否则会进行封号一周处理。
* content String 主体内容

返回值示例


{success: true, topic_id: 5433d5e4e737cbe96dcef312}


#### post /topics/update 编辑主题

接收 post 参数

* accesstoken String 用户的 accessToken
* topic_id String 主题id
* title String 标题
* tab String 目前有 ask share job
* content String 主体内容

返回值示例


{success: true, topic_id: '5433d5e4e737cbe96dcef312'}



### 主题收藏

#### post /topic_collect/collect 收藏主题

接收 post 参数

* accesstoken String 用户的 accessToken
* topic_id String 主题的id

返回值示例


{"success": true}


#### post /topic_collect/de_collect 取消主题

接收 post 参数

* accesstoken String 用户的 accessToken
* topic_id String 主题的id

返回值示例


{success: true}

#### get /topic_collect/:loginname 用户所收藏的主题

示例：[/api/v1/topic_collect/alsotang](/api/v1/topic_collect/alsotang)


### 评论

#### post /topic/:topic_id/replies 新建评论

接收 post 参数

* accesstoken String 用户的 accessToken
* content String 评论的主体
* reply_id String 如果这个评论是对另一个评论的回复，请务必带上此字段。这样前端就可以构建出评论线索图。

返回值示例


{success: true, reply_id: '5433d5e4e737cbe96dcef312'}


#### post /reply/:reply_id/ups 为评论点赞

接受 post 参数

* accesstoken String

接口会自动判断用户是否已点赞，如果否，则点赞；如果是，则取消点赞。点赞的动作反应在返回数据的 action 字段中，up or down。

返回值示例


{"success": true, "action": "down"}


### 用户

#### get /user/:loginname 用户详情

示例：[/api/v1/user/alsotang](/api/v1/user/alsotang)

#### post /accesstoken 验证 accessToken 的正确性

接收 post 参数

* accesstoken String 用户的 accessToken

如果成功匹配上用户，返回成功信息。否则 403。

返回值示例


{success: true, loginname: req.user.loginname, id: req.user.id, avatar_url: req.user.avatar_url}


### 消息通知

#### get /message/count 获取未读消息数

接收 get 参数

* accesstoken String

返回值示例


{ data: 3 }



#### get /messages 获取已读和未读消息

接收 get 参数

* accesstoken String
* mdrender String 当为 false 时，不渲染。默认为 true，渲染出现的所有 markdown 格式文本。

返回值示例


{
  data: {
    has_read_messages: [],
    hasnot_read_messages: [
      {
        id: "543fb7abae523bbc80412b26",
        type: "at",
        has_read: false,
        author: {
          loginname: "alsotang",
          avatar_url: "https://avatars.githubusercontent.com/u/1147375?v=2"
        },
        topic: {
          id: "542d6ecb9ecb3db94b2b3d0f",
          title: "adfadfadfasdf",
          last_reply_at: "2014-10-18T07:47:22.563Z"
        },
        reply: {
          id: "543fb7abae523bbc80412b24",
          content: "[@alsotang](/user/alsotang) 哈哈",
          ups: [ ],
          create_at: "2014-10-16T12:18:51.566Z"
          }
        },
    ...
    ]
  }
}


#### post /message/mark_all 标记全部已读

接收 post 参数

* accesstoken String

返回值示例


{ success: true,
  marked_msgs: [ { id: "544ce385aeaeb5931556c6f9" } ] }



#### post /message/mark_one/:msg_id 标记单个消息为已读

请求示例：[/message/mark_one/58ec7d39da8344a81eee0c14](/message/mark_one/58ec7d39da8344a81eee0c14)

接收 post 参数

* accesstoken String

返回值示例

js
{
  success: true,
  marked_msg_id: "58ec7d39da8344a81eee0c14"
}


### 知识点

1. 如何获取 accessToken？
    用户登录后，在设置页面可以看到自己的 accessToken。
    建议各移动端应用使用手机扫码的形式登录，验证使用 /accesstoken 接口，登录后长期保存 accessToken。	
`)
var userModel = new(models.UserModel)
var topicModel = new(models.TopicModel)
func Index(c *gin.Context) {
  //c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
  var no_reply_topics []models.Topic;

  var tops =[]string{"2","2"};
  tab:=c.Request.FormValue("tab")
  var queryTab string
  if tab==""{
    tab="all"
  }
  if tab=="all"{
    
  }else{
    queryTab=tab
  }
  var good bool
  if tab=="good"{
    good=true
  }
  log.Println(tab) 
  session := sessions.Get(c)
  var name string
  user:=models.User{}
  var err error
  log.Println(user)
  if nil!=session.Get("loginname"){
    name=session.Get("loginname").(string)
    user,err=userModel.GetUserByName(name)
  }
  log.Println(err)
  topics:=make([]models.Topic,10)

  log.Println(queryTab)
  log.Println(good)
  topics,_=topicModel.GetTopicByQuery(queryTab,good)
  log.Println(topics)
  base_url:="/?tab="+tab+"&page="
  var current_page int=1
  var pages int=1
  var page_start int
  var page_end int
  if (current_page-2)>0{
    page_start=current_page-2
  }else{
    page_start=1
  }
  if(page_start+4)>pages{
    page_end=pages
  }else{
    page_end=page_start+4
  }
  pagesArray:=[]string{}
  var i int
  for i =1;i<pages+1;i++{
    pagesArray=append(pagesArray,strconv.Itoa(i))
  }
  no_reply_topics2,err2:=cache.Get("no_reply_topics")
  json.Unmarshal(no_reply_topics2.([]byte),&no_reply_topics)
  log.Println("temp")
  log.Println(err2)
  //log.Println(temp)
  if(err!=nil){
    no_reply_topics,_=topicModel.GetTopicNoReply()
    no_reply_topics_json,_:=json.Marshal(no_reply_topics)
    cache.Set("no_reply_topics",no_reply_topics_json)
  }
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "布局页面",
    "no_reply_topics":no_reply_topics,
    "tops":tops,
    "user":user,
    "pagesArray":pagesArray,
    "base_url":base_url,
    "current_page":current_page,
    "topics":topics,
    "tab":tab,
    "pages":pages,
    "page_start":page_start,
    "page_end":page_end,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},
	})
}
var getstart =[]byte(`
## Node.js 入门

《**汇智网 Node.js 课程**》

http://www.hubwiz.com/course/?type=nodes

《**快速搭建 Node.js 开发环境以及加速 npm**》

http://fengmk2.com/blog/2014/03/node-env-and-faster-npm.html

《**Node.js 包教不包会**》

https://github.com/alsotang/node-lessons

《**ECMAScript 6入门**》

http://es6.ruanyifeng.com/

《**七天学会NodeJS**》

https://github.com/nqdeng/7-days-nodejs

《**Node入门-_一本全面的Node.js教程_**》

http://www.nodebeginner.org/index-zh-cn.html

## Node.js 资源

《**node weekly**》

http://nodeweekly.com/issues

《**node123-_node.js中文资料导航_**》

https://github.com/youyudehexie/node123

《**A curated list of delightful Node.js packages and resources**》

https://github.com/sindresorhus/awesome-nodejs

《**Node.js Books**》

https://github.com/pana/node-books

## Node.js 名人

《**名人堂**》

https://github.com/cnodejs/nodeclub/wiki/%E5%90%8D%E4%BA%BA%E5%A0%82

## Node.js 服务器

新手搭建 Node.js 服务器，推荐使用无需备案的 [DigitalOcean(https://www.digitalocean.com/)](https://www.digitalocean.com/?refcode=eba02656eeb3)		
`)
var about =[]byte(`
### 关于
CNode 社区为国内最大最具影响力的 Node.js 开源技术社区，致力于 Node.js 的技术研究。

CNode 社区由一批热爱 Node.js 技术的工程师发起，目前已经吸引了互联网各个公司的专业技术人员加入，我们非常欢迎更多对 Node.js 感兴趣的朋友。

CNode 的 SLA 保证是，一个9，即 90.000000%。

社区目前由 [@alsotang](http://cnodejs.org/user/alsotang) 在维护，有问题请联系：[https://github.com/alsotang](https://github.com/alsotang)

请关注我们的官方微博：http://weibo.com/cnodejs

### 移动客户端

客户端由 [@soliury](https://cnodejs.org/user/soliury) 开发维护。

源码地址： https://github.com/soliury/noder-react-native 。

立即体验 CNode 客户端，直接扫描页面右侧二维码。

另，安卓用户同时可选择：https://github.com/TakWolf/CNode-Material-Design ，这是 Java 原生开发的安卓客户端。		
`)
func About(c *gin.Context) {
  output := template.HTML(blackfriday.Run(about))
  session := sessions.Get(c)
  var name string
  user:=models.User{}
  //var err error
  log.Println(user)
  if nil!=session.Get("loginname"){
    name=session.Get("loginname").(string)
    user,_=userModel.GetUserByName(name)
  }
	c.HTML(http.StatusOK, "about", gin.H{
		"title": "布局页面",
    "about":    output,
    "user":user,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},   
	})
}

func Api(c *gin.Context) {

   output := template.HTML(blackfriday.Run(api))
   session := sessions.Get(c)
   var name string
   user:=models.User{}
   //var err error
   log.Println(user)
   if nil!=session.Get("loginname"){
     name=session.Get("loginname").(string)
     user,_=userModel.GetUserByName(name)
   } 
   c.HTML(http.StatusOK, "api", gin.H{
	   "title": "布局页面",
     "api":    output,
     "user":user,
     "config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},     
   })
}
func Getstart(c *gin.Context) {

  output := template.HTML(blackfriday.Run(getstart))
  session := sessions.Get(c)
  var name string
  user:=models.User{}
  //var err error
  log.Println(user)
  if nil!=session.Get("loginname"){
    name=session.Get("loginname").(string)
    user,_=userModel.GetUserByName(name)
  }

	c.HTML(http.StatusOK, "getstart", gin.H{
		"title": "布局页面",
    "getstart":    output,
    "user":user,
		"config": gin.H{
			"description": "CNode：Node.js专业中文社区",
		},    
	})
 }

