package main

import (
	"blog/helper"
	"blog/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (app *application) AddPostPage(ctx *gin.Context) {
	_, err := ctx.Cookie("session")
	if err != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}
	ctx.HTML(http.StatusOK, "add.html", models.Post{})
}

func (app *application) AddPost(ctx *gin.Context) {
	var post = &models.Post{}

	session, err := ctx.Cookie("session")
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}
	if helper.Length(ctx.PostForm("title")) || helper.Length(ctx.PostForm("message")) {
		ctx.Redirect(http.StatusFound, "/addpostpage")
		return
	}
	post.Id = uuid.New().String()
	post.Userid = session
	post.Title = ctx.PostForm("title")
	post.Message = ctx.PostForm("message")
	post.Like = 0

	err = app.blog.AddPostToDatabase(post)
	if err != nil {
		panic(err)
	}
	ctx.Redirect(http.StatusFound, "/")
}

func (app *application) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "edit.html", nil)
}

func (app *application) View(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "edit.html", nil)
}
