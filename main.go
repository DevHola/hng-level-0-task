package main

import (
	"time"
	"webservice/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main () {
	router := gin.Default()
	c:= cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "Delete"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,

	}
	router.Use(cors.New(c))
	router.GET("/", api.Index)
	router.Run(":8080")
}