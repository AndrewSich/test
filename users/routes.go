package users

import (
	"github.com/gin-gonic/gin"
)

func RouterUsers(route *gin.RouterGroup) {
	route.GET("/", FindAllUser)
	route.GET("/:id", FindUserByID)
	route.POST("/", CreateUser)
	// OPTIONS
	route.POST("/:id/add", UserAddContact)
	route.GET("/:id/list", UserListContact)
}
