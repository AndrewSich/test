package messages

import (
	"fmt"
	"time"

	"test/config"
)

type Message struct {
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

type MsgStore struct {
	Uid      string    `json:"uid"`
	Tid      string    `json:"tid"`
	Sender   bool      `json:"sender"`
	Status   string    `json:"status"`
	Type     string    `json:"type"`
	Data     string    `json:"data"`
	SendTime time.Time `json:"send-time"`
}

// component penting
// ID
// toID
// FromID
// Status
// Type
// Data
// ReceivedTime
// SendTime
// ReceiptServer

func CreateMessage(tid, fid, tipe, data string, timee time.Time) Message {
	db := config.GetDB()

	message := Message{
		ToID:          tid,
		FromID:        fid,
		Status:        "wait",
		Type:          tipe,
		Data:          data,
		ReceivedTime:  time.Now(),
		SendTime:      timee, //Use time from client
		ReceiptServer: time.Now(),
	}
	db.Model(&Message{}).Create(&message)

	return message
}

func FindAllMessages(fid, tid string) []MsgStore {
	db := config.GetDB()
	var messages []Message
	var msgStore []MsgStore

	//db.Model(&Message{}).Where("to_id IN ? AND from_id IN ?", []string{tid, fid}, []string{fid, tid}).Find(&messages)
	db.Model(&Message{}).Where("to_id IN (?) AND from_id IN (?)", []string{tid, fid}, []string{fid, tid}).Find(&messages)
	for _, message := range messages {
		sender := true
		if message.FromID == tid {
			sender = false
		}
		msg := MsgStore{
			Uid:      message.FromID,
			Tid:      message.ToID,
			Sender:   sender,
			Status:   message.Status,
			Type:     message.Type,
			Data:     message.Data,
			SendTime: message.SendTime,
		}

		msgStore = append(msgStore, msg)
	}

	return msgStore
}
