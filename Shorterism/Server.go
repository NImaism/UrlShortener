package main

import (
	database "Shorterism/Database"
	"Shorterism/Routing"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	// INIT SERVER
	Server := gin.Default()
	database.Connect()

	// SET LIMIT

	// SET ROUTING
	Routing.SetRouting(Server)

	// RUN SERVER
	if err := Server.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
