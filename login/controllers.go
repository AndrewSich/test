package login

import (
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

	db := config.GetDB()
	var user users.User

	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
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
