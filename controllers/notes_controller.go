package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-notes/models"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"notes/new.html",
		gin.H{},
		// helpers.SetPayload(c, gin.H{}),
	)
}

type FormData struct {
	Name    string `form:"name"`
	Content string `form:"content"`
}
