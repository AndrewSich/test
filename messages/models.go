package messages

import (
	"time"
)

type Message struct {
	ID            string    `gorm:"primary_key" json:"id"`
	ToID          string    `gorm:"column:to_id;not null" json:"to_id"`
	FromID        string    `gorm:"column:from_id;not null" json:"from_id"`
	Status        string    `gorm:"column:status" json:"status"`
	Type          string    `gorm:"column:type;not null" json:"type"`
	Data          string    `gorm:"column:data;size:2048;not null" json:"data"`
	MediaUrl      string    `gorm:"column:media_url" json:"media_url"`
	MediaSize     string    `gorm:"column:media_size" json:"media_size"`
	MediaName     string    `gorm:"column:media_name" json:"media_name"`
	ReceivedTime  time.Time `gorm:"column:received_time;default:CURRENT_TIMESTAMP" json:"received_time"`
	SendTime      time.Time `gorm:"column:send_time;default:CURRENT_TIMESTAMP" json:"send_time"`
	ReceiptServer time.Time `gorm:"column:receipt_server;default:CURRENT_TIMESTAMP" json:"receipt_server"`
}
