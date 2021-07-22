package messages

import "github.com/gin-gonic/gin"

func RouterMessages(route *gin.RouterGroup) {
	route.GET("/", MessagesPage)
	// Personal Messages
	route.GET("/personal", PersonalMessages)
	// route.POST("/personal", NewPersonalMessage)
	//route.GET("/group", GroupMessagesPage)
}
