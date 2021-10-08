package main

import (
	"blog/helper"
	"blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (app *application) AddPostPage (ctx *gin.Context){
	_, err := ctx.Cookie("session")
	if err != nil {
		ctx.Redirect(http.StatusFound,"/login")
		return
	}
	ctx.HTML(http.StatusOK,"add.html",nil)
}

func(app *application) AddPost (ctx *gin.Context){
	var post *models.Post
	var err error
	post.Userid, err = ctx.Cookie("session")
	if err != nil {
		ctx.Redirect(http.StatusFound,"/login")
		return
	}
	if helper.Length(ctx.PostForm("title")) || helper.Length(ctx.PostForm("message")){
		ctx.Redirect(http.StatusFound, "/addpostpage")
		return
	}
	post.Title = ctx.PostForm("title")
	post.Message = ctx.PostForm("message")
	post.Like = 0

	err = app.blog.AddPostToDatabase(post)
	if err != nil{
		panic(err)
	}
}
