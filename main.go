package main

import (
	"fmt"
	"net/http"

	"test/config"
	"test/users"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{})
}

func main() {
	db := config.TestDBInit()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gun.H{"data": "WELCOME TO REST API FLOME BY ANDREW SETYAWAN"})
	})

	rU := r.Group("/")
	users.RouterUsers(rU.Group("/users"))

	//Run Server
	fmt.Println("Server Running on==>FLOME Server by ANDREW SETYAWAN")
	r.Run()
}
