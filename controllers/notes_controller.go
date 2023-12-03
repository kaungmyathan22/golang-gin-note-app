package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaungmyathan22/golang-notes/controllers/helpers"
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

func NoteCreate(c *gin.Context) {
	name := c.PostForm("name")
	content := c.PostForm("content")
	models.NoteCreate(name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes")
}

type FormData struct {
	Name    string `form:"name"`
	Content string `form:"content"`
}

func NotesShow(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	// if currentUser == nil || currentUser.ID == 0 {
	// 	c.HTML(
	// 		http.StatusUnauthorized,
	// 		"notes/index.html",
	// 		helpers.SetPayload(c, gin.H{
	// 			"alert": "Unauthorized Access!",
	// 		}),
	// 	)
	// 	return
	// }
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(id)
	c.HTML(
		http.StatusOK,
		"notes/show.html",
		helpers.SetPayload(c, gin.H{
			"note": note,
		}),
	)
}

func NotesEditPage(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	// if currentUser == nil || currentUser.ID == 0 {
	// 	c.HTML(
	// 		http.StatusUnauthorized,
	// 		"notes/index.html",
	// 		helpers.SetPayload(c, gin.H{
	// 			"alert": "Unauthorized Access!",
	// 		}),
	// 	)
	// 	return
	// }
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(id)
	// note := models.NotesFind(currentUser, id)
	c.HTML(
		http.StatusOK,
		"notes/edit.html",
		helpers.SetPayload(c, gin.H{
			"note": note,
		}),
	)
}

func NotesUpdate(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	// if currentUser == nil || currentUser.ID == 0 {
	// 	c.HTML(
	// 		http.StatusUnauthorized,
	// 		"notes/index.html",
	// 		helpers.SetPayload(c, gin.H{
	// 			"alert": "Unauthorized Access!",
	// 		}),
	// 	)
	// 	return
	// }
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(id)
	// note := models.NotesFind(currentUser, id)
	name := c.PostForm("name")
	content := c.PostForm("content")
	note.Update(name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes/"+idStr)
}

func NotesDelete(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	// if currentUser == nil || currentUser.ID == 0 {
	// 	c.HTML(
	// 		http.StatusUnauthorized,
	// 		"notes/index.html",
	// 		helpers.SetPayload(c, gin.H{
	// 			"alert": "Unauthorized Access!",
	// 		}),
	// 	)
	// 	return
	// }
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.NotesMarkDelete(id)
	c.Redirect(http.StatusSeeOther, "/notes")
}
