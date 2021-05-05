package main

import (
	"log"
	"net/http"
	"os"

	// "github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//mux routing
	http.HandleFunc("/", handlerIndex)

	// run server
	log.Printf("Server Running On\n")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println(err.Error())
	}

	// router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

	// router.Run(":" + port)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Levxa Web Server"))
}
