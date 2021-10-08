package main

import (
	"blog/helper"
	_"blog/helper"
	"blog/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (app *application) Home(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"home.html",models.BlogUser{})
}

func (app *application) SignUpPage(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"signup.html",nil)
}

func (app *application) Signup(ctx *gin.Context){
	log.Println("Working first")
	b := models.BlogUser{}
	name := ctx.PostForm("name")
	password := helper.HashPassword(ctx.PostForm("password"))
	check := helper.IsValidEmail(ctx.PostForm("email"))

	if  helper.Length(name) || helper.Length(password) || helper.Length(ctx.PostForm("email")) {
		return
	}
	if check != true {
		log.Println("Invalid Email")
		return
	}

	b.Id = uuid.New().String()
	b.Name =  name
	b.PassWord = password
	b.Email = ctx.PostForm("email")
	b.Bio = ""
	b.State = true
	b.Followers = 0

	log.Println("working middle")

	err := app.user.QueryUser(&b,b.Email)
	if err != nil {
		err = app.user.SignUpUser(&b)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
	log.Println("Still Testing")

	log.Println("working last")
}

func (app *application) LoginPage(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"login.html",models.BlogUser{})
}


func (app *application) Login(ctx *gin.Context){

	log.Println("working first")
	password :=ctx.PostForm("password")
	check := helper.IsValidEmail(ctx.PostForm("email"))
	if  helper.Length(password) || helper.Length(ctx.PostForm("email")) {
		return
	}
	if check != true {
		return
	}
	log.Println("working second")
	b,err := app.user.QueryEmail(ctx.PostForm("email"))
	if err != nil {
		panic(err)
	}
	log.Println("working after panic")
	ctx.SetCookie("session",b.Id,3600,"/","localhost",true,true)
	ok := helper.ComparePassword(b.PassWord, password)
	if ok{
		ctx.Redirect(http.StatusFound,"/")
		log.Println("working after setting cookie")
		return
	}
	ctx.String(http.StatusNotFound,"could not login in")

}





