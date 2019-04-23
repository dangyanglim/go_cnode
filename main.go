package main

import (
	db "github.com/dangyanglim/go_cnode/database"
	//"log"
)

func main() {
	defer db.SqlDB.Close()
	defer db.MogSession.Close()
	defer db.Redis.Close()

	router := initRouter()

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AllowCredentials = true
	// config.AllowMethods = []string{"*"}
	// config.AllowHeaders = []string{"*"}
	//router.Use(cors.Default())
	router.Run(":9035")
}
