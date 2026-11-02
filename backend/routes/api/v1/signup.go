package v1

import (
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/TwoKrayyz/krayyzchat/database"
)

type SignUpForm struct {
	Username	string	`json:"username" form:"username" binding:"required"`
	Password	string	`json:"password" form:"password" binding:"required"`
}

func SignUp(c *gin.Context) {
	var signUpData SignUpForm

	if err := c.ShouldBind(&signUpData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}

	username := signUpData.Username
	password := signUpData.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}

	db := database.DB

	var dbUsername string

	error := db.QueryRow("SELECT username FROM users WHERE username = $1", username).Scan(&dbUsername)

	if error == sql.ErrNoRows {
		_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"data": "User created!"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	} else {
		c.JSON(http.StatusFound, gin.H{"data": "User is already in the database!"})
	}
}