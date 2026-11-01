package routes

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.File("../frontend/public/index.html")
}