package routes

import "github.com/gin-gonic/gin"

func SignIn(c *gin.Context) {
	c.File("../frontend/public/pages/sign-in.html")
}