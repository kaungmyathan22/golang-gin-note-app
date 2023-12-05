package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-notes/controllers"
	controller_helpers "github.com/kaungmyathan22/golang-notes/controllers/helpers"
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
	r.GET("/notes/:id", controllers.NotesShow)
	r.POST("/notes", controllers.NoteCreate)
	r.GET("/notes/edit/:id", controllers.NotesEditPage)
	r.POST("/notes/:id", controllers.NotesUpdate)
	r.DELETE("/notes/:id", controllers.NotesUpdate)

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", controller_helpers.SetPayload(c, gin.H{
			"title":     "Notes application",
			"logged_in": controller_helpers.IsUserLoggedIn(c),
		}))
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "Notes applications",
		})
	})
	log.Println("Server Started!")
	r.Run()
}
