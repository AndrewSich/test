package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	// "github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type Mux struct {
	blog, main *http.ServeMux
}

func (mux Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("[+] success ", r.Host)
	if r.Host == "levxa-web-server.herokuapp.com" {
		mux.main.ServeHTTP(w, r)
		return
	}

	// subdomain checking
	domainParts := strings.Split(r.Host, ".")
	if domainParts[0] == "blog" {
		mux.blog.ServeHTTP(w, r)
	} else {
		log.Printf("[!] warning Subdomain is not found!")
		http.Error(w, "Pages Not Found", 404)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//mux routing
	mux := &Mux{
		blog: http.NewServeMux(),
		main: http.NewServeMux(),
	}

	mux.blog.HandleFunc("/", handlerBlog)
	mux.main.HandleFunc("/", handlerIndex)

	// run server

	log.Printf("Server Running On\n")
	err := http.ListenAndServe(":"+port, mux)
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

func handlerBlog(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This a Blog Pages", r.URL.Path[1:])
	w.Write([]byte("This a Blog Pages"))
}
