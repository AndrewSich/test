package users

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"

	"test/chats"
	"test/config"
	"test/messages"
)

type FormUser struct {
	Id           string `json:"id"`
	Nickname     string `json:"nickname"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Address      string `json:"address"`
	ProfileImage string `json:"profile_image"`
	BannerImage  string `json:"banner_image"`
	Bio          string `json"bio"`
}

type FormMessage struct {
	Type     string    `json:"type"`
	Data     string    `json:"data"`
	SendTime time.Time `json:"send_time"`
}

// Find All User
func FindAllUser(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	users := GetAllUser()
	c.JSON(200, users)
}

// Find User By ID
func FindUserByID(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	uid := c.Param("id")
	user := GetUserByID(uid)

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

// ========================== CONTACTS =======================================
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

// ================================= CHATS ==================================
func UserAddChat(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	db := config.GetDB()
	var profile Profile
	var form FormUser
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println("[FLOME] => error: ", err.Error())
		c.JSON(400, gin.H{"data": err.Error()})
		return
	}

	parentID := c.Param("id")
	childID := form.Id

	db.Model(&Profile{}).Where("id = ?", childID).Take(&profile)
	data := chats.Chat{
		ParentID: parentID,
		ChildID:  profile.ID,
		Nickname: profile.Nickname,
		Image:    profile.ProfileImage,
	}

	db.Model(&chats.Chat{}).Create(&data)
	file := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	c.JSON(200, file)
}

func UserListChat(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	uid := c.Param("id")
	data := chats.FindAllChat(uid)

	c.JSON(200, data)
}

// ========================== MESSAGES ================================
func UserAddMessage(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	db := config.GetDB()

	var profile Profile
	var form FormMessage
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println("[FLOME] => error: ", err.Error())
		c.JSON(400, gin.H{"data": err.Error()})
		return
	}

	fid := c.Param("id")
	tid := c.Param("tid")
	tipe := form.Type
	data := form.Data
	timee := form.SendTime

	// Add Chat
	exist := chats.CheckExist(fid, tid)
	if exist == false {
		db.Model(&Profile{}).Where("id = ?", tid).Take(&profile)
		chat := chats.Chat{
			ParentID:          fid,
			ChildID:           tid,
			Nickname:          profile.Nickname,
			Image:             profile.ProfileImage,
			LastMessage:       "Halloo",
			LastMessageStatus: "wait",
			LastMessageTime:   time.Now(),
		}

		db.Model(&chats.Chat{}).Create(&chat)
	}

	message := messages.CreateMessage(tid, fid, tipe, data, timee)
	c.JSON(200, message)
}

func UserListMessage(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	fid := c.Param("id")
	tid := c.Param("tid")
	data := messages.FindAllMessages(fid, tid)

	c.JSON(200, data)
}
