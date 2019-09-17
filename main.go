package main

import (
	db "go_cnode/database"
	"encoding/json"
	"os"
	"go_cnode/router"
	//"log"
)
type configuration struct {
    Port string
	Mongo_url   string
	Redis_url	string
}
func LoadConf()(conf configuration ){
    // 打开文件
    file, _ := os.Open("conf.json")

    // 关闭文件
    defer file.Close()

    //NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
    decoder := json.NewDecoder(file)


    //Decode从输入流读取下一个json编码值并保存在v指向的值里
	decoder.Decode(&conf)
	return conf
}

func main() {
	conf:=LoadConf()
	db.Config(conf.Mongo_url,conf.Redis_url)
	//defer db.SqlDB.Close()
	defer db.MogSession.Close()
	defer db.Redis.Close()

	router := router.InitRouter()

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AllowCredentials = true
	// config.AllowMethods = []string{"*"}
	// config.AllowHeaders = []string{"*"}
	//router.Use(cors.Default())
	router.Run(conf.Port)
}
