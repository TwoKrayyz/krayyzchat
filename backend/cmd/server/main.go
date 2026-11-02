package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/TwoKrayyz/krayyzchat/routes"
	"github.com/TwoKrayyz/krayyzchat/database"
	"github.com/TwoKrayyz/krayyzchat/routes/api/v1"
)

func main() {
	db := database.DB

	if err := database.Pool(); err != nil {
		log.Fatal("Database pool failed:", err)
	}
	defer db.Close()

	gin.SetMode(gin.DebugMode)

	router := gin.New()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Static("/public", "../frontend/public")
	router.GET("/", routes.Index)
	router.GET("/signin", routes.SignIn)
	router.GET("/signup", routes.SignUp)

	v1Group := router.Group("/api/v1")
	{
		v1Group.POST("/signup", v1.SignUp)
	}

	router.Run(":3000")
}