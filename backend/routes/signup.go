package routes

import "github.com/gin-gonic/gin"

func SignUp(c *gin.Context) {
	c.File("../frontend/public/pages/sign-up.html")
}