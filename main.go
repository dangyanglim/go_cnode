package main

import (
	db "go_cnode/database"
	"go_cnode/router"
	"go_cnode/utils"
	//"log"
)



func main() {
	conf := utils.LoadConf()
	db.Config(conf.Mongo_url, conf.Redis_url)
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
