package main

import (
	"fmt"

	"test/config"
	"test/login"
	"test/messages"
	"test/users"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"

	static "github.com/gin-contrib/static"
	"gopkg.in/olahol/melody.v1"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&messages.Message{})
}

func main() {
	db := config.TestDBInit()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	m := melody.New()

	// Static file route
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
	// WEB SOCKET ENDOINT
	r.GET("/socket", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	rU := r.Group("/")
	rL := r.Group("/")
	rM := r.Group("/")

	// Login
	login.RouterLogin(rL.Group("/login"))
	// Users
	users.RouterUsers(rU.Group("/users"))
	// Messages
	messages.RouterMessages(rM.Group("/messages"))

	//Run Server
	fmt.Println("Server Running on==>FLOME Server by ANDREW SETYAWAN")
	r.Run()
}
