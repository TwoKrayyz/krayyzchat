package main

import (
	"github.com/gin-gonic/gin"
	"github.com/TwoKrayyz/krayyzchat/routes"
)

func main() {
	gin.SetMode(gin.DebugMode)

	router := gin.New()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Static("/public", "../frontend/public")
	router.GET("/", routes.Index)
	router.GET("/signup", routes.SignUp)

	router.Run(":3000")
}