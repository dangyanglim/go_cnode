# go_cnode

#### 项目介绍
http://cnodejs.org/ 是一个用nodejs语言编写的开源论坛  
打算是用golang语言仿写  
项目部署在 http://fenghuangyu.cn:9035 账号密码admin   
#### 安装
go version 1.12.4  
git version 2.19  
go版本>1.11  
git版本>2.17  
用了go mod 管理依赖  
安装mongodb  
安装redis  

## 使用
 
### 使用命令行

```bash
$ git clone https://github.com/dangyanglim/go_cnode.git
$ cd go_cnode
$ redis-server                     # 要安装redis
$ go run main.go                   # 访问 http://localhost:9035
```
#### 功能介绍
- 邮箱注册/Github第三方注册  
- Go 模块管理  
- 后台 Gin+mongodb+redis
- 前台 bootstrap+jquery+渲染模板  

![go.png](go.png)  

## 目录结构  
```
├─.vscode
├─controllers
│  ├─reply
│  ├─sign
│  ├─site
│  └─topic
├─database
├─mgoModels
├─public
│  ├─images
│  ├─img
│  ├─javascripts
│  ├─libs
│  │  ├─bootstrap
│  │  │  ├─css
│  │  │  ├─img
│  │  │  └─js
│  │  ├─code-prettify
│  │  ├─editor
│  │  │  └─fonts
│  │  ├─font-awesome
│  │  │  ├─css
│  │  │  └─fonts
│  │  └─webuploader
│  ├─stylesheets
│  └─upload
├─router
├─service
│  ├─cache
│  └─mail
├─utils
└─views
    ├─about
    ├─api
    ├─common
    ├─edit
    ├─getStart
    ├─index
    ├─message
    ├─notify
    ├─searchPass
    ├─setting
    ├─showSignUp
    ├─signIn
    ├─signUp
    └─topic
    
```
 
  
 

