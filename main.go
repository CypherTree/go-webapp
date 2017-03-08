package main

import (
	"fmt"
	"net/http"

	"go-webapp/config"
	"go-webapp/db"
)

func main() {
	fmt.Println("Init server...")

	// configure environment variables
	config.Settings.MakeConfig()
	fmt.Println("Port: ", config.Settings.Port)

	http.Handle("/", http.FileServer(http.Dir("public")))
	router := registerHandlers()
	fmt.Println("Registered handlers...")

	//Start dbs
	db.Conn = db.MakeConn(config.Settings.DbURL, config.Settings.DbName)
	db.Redis = db.MakeRedisConn(config.Settings.RedisURL, config.Settings.RedisPwd)

	router.Run(":" + config.Settings.Port)
}
