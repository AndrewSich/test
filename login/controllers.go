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
	db := config.GetDB()
	var form FormLogin
	var user users.User

	body := c.Request.Body
	d, _ := ioutil.ReadAll(body)
	fmt.Println(body)
	fmt.Println(string(d))

	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println("[FLOME] => error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	username := form.Username
	password := form.Password
	db.Model(&users.User{}).Where("username = ? AND password = ?", username, password).Find(&user)
	if user.Username == username && user.Password == password {

		// data := map[string]interface{}{
		// 	"id":           user.ID,
		// 	"nickname":     user.Nickname,
		// 	"username":     user.Username,
		// 	"email":        user.Email,
		// 	"address":      user.Address,
		// 	"profil_image": user.ProfileImage,
		// 	"banner_image": user.BannerImage,
		// 	"bio":          user.Bio,
		// }

		c.JSON(http.StatusOK, gin.H{"data": "Congrats, Success login.."})
	} else {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or Password Not Found"})
	}

}
