package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Exam andersen"})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("index.html")

	router.GET("/", HelloWorld)
	err := router.Run(":80")
	if err != nil {
		fmt.Print("all ok")
	}

}