package users

import (
	"github.com/gin-gonic/gin"
)

func RouterUsers(route *gin.RouterGroup) {
	route.GET("/", FindAllUser)
	route.GET("/:id", FindUserByID)
	route.POST("/", CreateUser)
	// route contacts
	route.POST("/:id/add", UserAddContact)
	route.GET("/:id/list", UserListContact)
	// route Chats
	route.POST("/:id/chats/add", UserAddChat)  // "/:id/chats" type POST
	route.GET("/:id/chats/list", UserListChat) // "/:id/chats" type GET
	// route message
	route.POST("/:id/chats/:tid", UserAddMessage)
	route.GET("/:id/chats/:tid", UserListMessage)
}
