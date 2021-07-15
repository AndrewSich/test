package users

import (
	"github.com/gin-gonic/gin"
)

func RouterUsers(route *gin.RouterGroup) {
	route.GET("/", FindAllUser)
	route.POST("/", CreateUser)
}
