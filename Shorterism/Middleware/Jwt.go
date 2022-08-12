package Middleware

import (
	"Shorterism/Model"
	"Shorterism/Utility"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := Utility.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, Model.ErrorResponse("Token Is Expire Or Incorrect", "Unauthorized"))
			c.Abort()
			return
		}
		c.Next()
	}
}
