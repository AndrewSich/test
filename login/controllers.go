package login

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"test/config"
	"test/users"

	"github.com/gin-gonic/gin"
)

type FormLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password"form:"password"`
}

func Login(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	db := config.GetDB()
	var form FormLogin
	var user users.User

	username := c.Param("username")
	password := c.Param("password")
	db.Model(&users.User{}).Where("username = ? AND password = ?", username, password).Find(&user)
	if user.Username == username && user.Password == password {

		c.JSON(200, user)

	} else {

		c.JSON(400, gin.H{"data": "Username or Password Not Found"})
	}

}
