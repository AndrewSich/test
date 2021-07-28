package users

import (
	"time"

	"test/config"
)

type User struct {
	ID           string    `gorm:"primary_key" json:"id"`
	Nickname     string    `gorm:"column:nickname;not null" json:"nickname"`
	Username     string    `gorm:"column:username;unique_index;not null" json:"username"`
	Email        string    `gorm:"column:email;unique_index;not null" json:"email"`
	Password     string    `gorm:"column:password;not null" json:"password"`
	Address      string    `gorm:"culomn:address" json:"address"`
	ProfileImage string    `gorm:"column:profile_image" json:"profile_image"`
	BannerImage  string    `gorm:"column:banner_image" json:"banner_image"`
	Bio          string    `gorm:"column:bio;size:1024" json"bio"`
	Contacts     []Profile `gorm:"many2many:user_contacts" json:"contacts"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Profile struct {
	ID           string    `gorm:"primary_key" json: "id"`
	Nickname     string    `gorm:"column:nickname" json:"nickname"`
	Username     string    `gorm:"column:username" json:"username"`
	Address      string    `gorm:"column:address" json:"address"`
	ProfileImage string    `gorm:"column:profile_image" json:"profile_image"`
	BannerImage  string    `gorm:"column:banner_image" json:"banner_image"`
	Bio          string    `gorm:"column:bio" json:"bio"`
	LastSeen     time.Time `gorm:"column:last_seen;default:CURRENT_TIMESTAMP" json:"last_seen"`
}

func GetAllUser() []User {
	db := config.GetDB()
	var users []User

	db.Model(&User{}).Find(&users)
	return users
}

func GetUserByID(uid string) User {
	db := config.GetDB()
	var user User

	db.Model(&User{}).Where("id = ?", uid).Take(&user)
	return user
}
