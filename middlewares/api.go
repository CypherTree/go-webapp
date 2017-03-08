package middlewares

import (
	"net/http"

	"go-webapp/db"

	"github.com/gin-gonic/gin"
)

// Cors - Cross Origin Resource Sharing middleware, all everyone
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Connect - makes the cloned `db` object available for each handler
func Connect(c *gin.Context) {
	s := db.Conn.Session.Clone()

	defer s.Close()

	c.Set("db", s)
	c.Next()
}

// ErrorHandler - middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": c.Errors,
		})
	}
}
