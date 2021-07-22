package messages

import (
	"time"

	"github.com/gin-gonic/gin"
)

type NewMessage struct {
	ToID      string    `json:"to_id"`
	FromID    string    `json:"from_id`
	Type      string    `json:"type"`
	Data      string    `json:"data"`
	MediaUrl  string    `json:"media_url"`
	MediaSize string    `json:"media_size"`
	MediaName string    `json:"media_name"`
	SendTime  time.Time `json:"send_time"`
}

func MessagesPage(c *gin.Context) {

	c.JSON(200, gin.H{"data": "This Flome Messages Bulk Page"})
}

func PersonalMessages(c *gin.Context) {
	// CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	c.JSON(200, gin.H{"data": "success"})
}

// func NewPersonalMessage(c *gin.Context) {
// 	// CORS
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

// 	db := config.getDB()
// 	var msg Message

// 	var newMsg NewMessage
// 	if err := c.ShouldBindJSON(&newMsg); err != nil {
// 		fmt.Println("[FLOME] => error: ", err.Error())
// 		c.JSON(400, gin.H{"data": err.Error()})
// 		return
// 	}

// 	msg := Message{
// 		ID:            "0",
// 		ToID:          newMsg.ToID,
// 		FromID:        newMsg.FromID,
// 		Status:        "wait",
// 		Type:          newMsg.Type,
// 		Data:          newMsg.Data,
// 		MediaUrl:      newMsg.MediaUrl,
// 		MediaSize:     newMsg.MediaSize,
// 		MediaName:     newMsg.MediaName,
// 		ReceivedTime:  time.Now(),
// 		SendTime:      newMsg.SendTime,
// 		ReceiptServer: time.Now(),
// 	}

// 	db.Model(&Message{}).Create(&msg)

// 	// Response Sender
// 	responseSender := map[string]interface{}{
// 		"target": newMsg.ToID,
// 		"sender": true,
// 		"status": "wait",
// 		"type":   newMsg.Type,
// 		"data":   newMsg.Data,
// 		"time":   newMsg.SendTime,
// 	}

// 	// Response Receiver
// 	responseReceiver := map[string]interface{}{
// 		"target": newMsg.FromID,
// 		"sender": false,
// 		"status": "wait",
// 		"type":   newMsg.Type,
// 		"data":   newMsg.Data,
// 		"time":   newMsg.SendTime,
// 	}
// }
