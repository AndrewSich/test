package users

import (
	"time"
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
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
