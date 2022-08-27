package main

import (
	"Shorterism/internal/database"
	"Shorterism/internal/handler"
	"Shorterism/internal/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"os"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}

	server := gin.Default()

	mongodb := database.New()

	hnd := handler.App{
		Store:  store.New(mongodb),
		Logger: logger.Named("Handler.App"),
	}

	hnd.SetRouting(server)

	// RUN SERVER
	if err := server.Run(os.Args[1]); err != nil {
		logger.Fatal("Can't run server", zap.Error(err))
	}
}
