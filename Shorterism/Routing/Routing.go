package Routing

import (
	"Shorterism/Controllers"
	"Shorterism/Middleware"
	"github.com/gin-gonic/gin"
)

func SetRouting(e *gin.Engine) {
	e.GET("/:link", Controllers.Link)

	Api := e.Group("/api/v1/").Use(Middleware.JwtAuthMiddleware())
	Api.POST("CreateLink/", Controllers.RegisterLink)

	Account := e.Group("/account/v1/")
	Account.POST("Register/", Controllers.Register)
	Account.POST("Login/", Controllers.Login)
}
