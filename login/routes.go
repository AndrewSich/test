package login

import (
	"github.com/gin-gonic/gin"
)

func RouterLogin(route *gin.RouterGroup) {
	// route user login
	route.POST("/", Login)
}
