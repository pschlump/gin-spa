package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	spa "github.com/pschlump/gin-spa/spa"
)

func main() {
	// Setup Test Code
	os.MkdirAll("./www", 0755)
	ioutil.WriteFile("./www/index.html", []byte("File: index.html\n"), 0644)
	// Run Server
	r := gin.Default()
	pathsToRemap := spa.SPAList{
		FromTo: map[string]string{
			"/login":  "/index.html",
			"/logout": "/index.html",
			"/home":   "/index.html",
		},
	}
	r.Use(spa.StaticServeMiddleware("/", "./www", pathsToRemap, os.Stderr)) // your build of React or other SPA
	fmt.Printf("Listening on http://127.0.0.1:9094/")
	r.Run("127.0.0.1:9094")
}
