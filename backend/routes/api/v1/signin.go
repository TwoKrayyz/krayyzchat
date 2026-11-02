package v1

import (
	"errors"
	"net/http"
	"database/sql"
	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/TwoKrayyz/krayyzchat/database"
)

type SignInForm struct {
	Username	string	`json:"username" form:"username" binding:"required"`
	Password	string	`json:"password" form:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	var signInData SignInForm

	if err := c.ShouldBind(&signInData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}

	username := signInData.Username
	password := signInData.Password

	db := database.DB

	var dbUsername string
	var dbHashedPassword string

	err := db.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&dbUsername, &dbHashedPassword)

	if errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, gin.H{"data": "User does not exist"})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(password))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"data": "Password or Username is incorrect"})
		} else {
			c.JSON(http.StatusFound, gin.H{"data": "Logged in!"})
		}
	}
}