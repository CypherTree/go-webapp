package main

import (
	"github.com/gin-gonic/gin"

	"go-webapp/handlers"
	"go-webapp/middlewares"
)

func registerHandlers() *gin.Engine {
	r := gin.Default()

	// enable cors
	r.Use(middlewares.Cors())

	// attach auth middleware
	r.Use(middlewares.AuthSetUser())

	// auth routes
	r.POST("/auth/facebook", handlers.Facebook)
	r.POST("/auth/instagram", handlers.Instagram)

	return r
}
