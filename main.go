package main

import (
	db "github.com/dangyanglim/go_cnode/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()

	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// config.AllowCredentials = true
	// config.AllowMethods = []string{"*"}
	// config.AllowHeaders = []string{"*"}
	//router.Use(cors.Default())
	router.Run(":9031")
}
