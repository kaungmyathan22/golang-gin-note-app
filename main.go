package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-notes/controllers"
	"github.com/kaungmyathan22/golang-notes/models"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Static("/vendor", "./static/vendors")
	models.ConnectDatabase()
	models.DBMigrate()
	log.Println("Successfully connected to database!")
	r.LoadHTMLGlob("templates/**/**")
	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes applications",
		})
	})
	log.Println("Server Started!")
	r.Run()
}
