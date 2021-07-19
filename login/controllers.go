package login

import (
	"fmt"
	"io/ioutil"

	"test/config"
	"test/users"

	"github.com/gin-gonic/gin"
)

// type FormLogin struct {
// 	Username string `json:"username" form:"username"`
// 	Password string `json:"password"form:"password"`
// }

func Login(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	body := c.Request.Body
	b, _ := ioutil.ReadAll(body)
	fmt.Println(string(b))

	db := config.GetDB()
	var user users.User

	username := c.Param("username")
	password := c.Param("password")
	if username == "" || password == "" {

		c.JSON(400, gin.H{"data": "username or password invalid"})
	} else {
		db.Model(&users.User{}).Where("username = ? AND password = ?", username, password).Find(&user)
		if user.Username == username && user.Password == password {

			c.JSON(200, user)

		} else {
			c.JSON(400, gin.H{"data": "username or password not found!!"})
		}
	}

}
