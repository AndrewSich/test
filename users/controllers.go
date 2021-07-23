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
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	//c.Header("Access-Control-Allow-Credentials", "true")
	// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	db := config.GetDB()
	var users []User

	db.Model(&User{}).Find(&users)
	c.JSON(200, users)
}

// Find User By ID
func FindUserByID(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	db := config.GetDB()
	var user User

	uid := c.Param("id")
	db.Model(&User{}).Where("id = ?", uid).Take(&user)

	c.JSON(200, user)
}

// Create a New User
func CreateUser(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

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
	profile := Profile{
		ID:           user.ID,
		Nickname:     form.Nickname,
		Username:     form.Username,
		Address:      form.Address,
		ProfileImage: form.ProfileImage,
		BannerImage:  form.BannerImage,
		Bio:          form.Bio,
		LastSeen:     time.Now(),
	}

	db.Model(&User{}).Create(&user)
	db.Model(&Profile{}).Create(&profile)
	data := map[string]interface{}{
		"user":    user,
		"profile": profile,
	}

	c.JSON(http.StatusOK, data)
}

// Add Contact
func UserAddContact(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	db := config.GetDB()
	var contact Profile
	var form FormUser

	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println("[FLOME] => error: ", err.Error())
		c.JSON(400, gin.H{"data": err.Error()})
		return
	}

	username := form.Username
	db.Model(&Profile{}).Where("username = ?", username).Take(&contact)

	uid := c.Param("id")
	var user User
	db.Model(&User{}).Where("id = ?", uid).Take(&user)

	user.Contacts = []Profile{contact}
	db.Omit("Contacts.*").Save(&user)
	//db.Model(&user).Association("Contacts").Append(&contact)

	c.JSON(200, gin.H{"data": "success add contact"})
}

// List Contact
func UserListContact(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	db := config.GetDB()
	var contacts []Profile
	var user User

	uid := c.Param("id")
	db.Model(&User{}).Where("id = ?", uid).Take(&user)

	db.Model(&user).Association("Contacts").Find(&contacts)
	c.JSON(200, contacts)
}
