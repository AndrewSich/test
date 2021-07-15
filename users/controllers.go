package users

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"test/config"
)

type FormUser struct {
	Nickname     string `json:"nickname"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Address      string `json:"address"`
	ProfileImage string `json:"profile_image"`
	BannerImage  string `json:"banner_image"`
	Bio          string `json"bio"`
}

// Find All User
func FindAllUser(c *gin.Context) {
	db := config.GetDB()
	var users []User

	db.Model(&User{}).Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	db := config.GetDB()
	var form FormUser
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println("[FLOME] => error: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	uid := uuid.New()
	user := User{
		ID:           uid.String(),
		Nickname:     form.Nickname,
		Username:     form.Username,
		Email:        form.Email,
		Password:     form.Password,
		Address:      form.Address,
		ProfileImage: form.ProfileImage,
		BannerImage:  form.BannerImage,
		Bio:          form.Bio,
		CreatedAt:    time.Now(),
	}

	db.Model(&User{}).Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
