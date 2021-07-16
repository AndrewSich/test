package main

import (
	"fmt"
	"net/http"

	"test/config"
	"test/login"
	"test/messages"
	"test/users"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"
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
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "WELCOME TO REST API FLOME BY ANDREW SETYAWAN"})
	})

	rU := r.Group("/")
	rL := r.Group("/")

	// Login
	login.RouterLogin(rL.Group("/login"))
	// Users
	users.RouterUsers(rU.Group("/users"))

	//Run Server
	fmt.Println("Server Running on==>FLOME Server by ANDREW SETYAWAN")
	r.Run()
}
