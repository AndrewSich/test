package chats

import (
	"fmt"
	"time"

	"test/config"
)

type Chat struct {
	ParentID          string    `gorm:"column:parent_id" json:"parent-id"`
	ChildID           string    `gorm:"child_id" json:"child-id"`
	Nickname          string    `gorm:"column:nickname" json:"nickname"`
	Image             string    `gorm:"column:image" json:"image"`
	LastMessage       string    `gorm:"column:last_message" json:"last-message"`
	LastMessageStatus string    `gorm:"column:last_message_status" json:"last-message-status"`
	LastMessageTime   time.Time `gorm:"column:last_message_time;default:CURRENT_TIMESTAMP" json:"last-message-time"`
}

func FindAllChat(uid string) []Chat {
	db := config.GetDB()
	var chats []Chat

	db.Model(&Chat{}).Where("parent_id = ?", uid).Find(&chats)
	return chats
}

func CheckExist(pid, cid string) bool {
	db := config.GetDB()
	var chat Chat

	db.Model(&Chat{}).Where("parent_id = ? AND child_id = ?", pid, cid).Take(&chat)
	exist := true
	if chat.ParentID == "" {
		fmt.Println("[FLOME] => Chat not found!")
		exist = false
	}

	return exist
}
