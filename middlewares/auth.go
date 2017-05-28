package middlewares

import (
	"go-webapp/models"

	"go-webapp/config"

	"go-webapp/dbapi/user"

	"github.com/gin-gonic/gin"
)

// AuthSetUser - Set user in context using auth token
func AuthSetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get(config.AuthHeaderName)
		session := &models.Session{}
		userapi.ValidateAuthToken(tokenString, session)

		c.Set("UserSession", session)
		c.Next()
	}
}

// IsAuthenticated - check whether user is authenticated or not
func IsAuthenticated() gin.HandlerFunc {

	return func(c *gin.Context) {
		if _, ok := c.Get("UserSession"); ok {
			c.Next()
		} else {
			c.AbortWithStatus(401)
		}
	}
}
