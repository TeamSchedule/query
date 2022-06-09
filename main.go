package main

import (
	"log"
	"query/server"
	"query/server/conf"
)

func main() {
	router := server.SetupServer()

	err := router.Run(":" + conf.ServerPort)
	if err != nil {
		log.Fatalln("Query service crashed on startup!")
		return
	}
}
