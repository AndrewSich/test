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
	route.POST("/:id/chats/add", UserAddChat)
	route.GET("/:id/chats/list", UserListChat)
}
