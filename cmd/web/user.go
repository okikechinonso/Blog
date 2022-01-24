package main

import (
	"blog/helper"
	_ "blog/helper"
	"blog/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (app *application) Home(c *gin.Context) {
	_, err := c.Cookie("session")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.HTML(http.StatusOK, ".html", models.Post{})
}

func (app *application) SignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func (app *application) Signup(c *gin.Context) {
	log.Println("Working first")
	b := models.BlogUser{}
	name := c.PostForm("name")
	password := helper.HashPassword(c.PostForm("password"))
	check := helper.IsValidEmail(c.PostForm("email"))

	if helper.Length(name) || helper.Length(password) || helper.Length(c.PostForm("email")) {
		return
	}
	if check != true {
		log.Println("Invalid Email")
		return
	}

	b.Id = uuid.New().String()
	b.Name = name
	b.PassWord = password
	b.Email = c.PostForm("email")
	b.Bio = ""
	b.State = true
	b.Followers = 0

	log.Println("working middle")

	err := app.user.QueryUser(&b, b.Email)
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

func (app *application) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", models.BlogUser{})
}

func (app *application) blogHome(c *gin.Context) {
	c.HTML(http.StatusOK, "bloghome.html", models.BlogUser{})
}

func (app *application) Login(c *gin.Context) {

	password := c.PostForm("password")
	check := helper.IsValidEmail(c.PostForm("email"))
	if helper.Length(password) || helper.Length(c.PostForm("email")) {
		return
	}
	if check != true {
		return
	}

	b, err := app.user.QueryEmail(c.PostForm("email"))
	if err != nil {
		panic(err)
	}
	log.Println("working after panic")
	c.SetCookie("session", b.Id, 3600, "/", "localhost", true, true)
	ok := helper.ComparePassword(b.PassWord, password)
	if ok {
		c.Redirect(http.StatusFound, "/blogHome")
		log.Println("working after setting cookie")
		return
	}
	c.String(http.StatusNotFound, "could not login in")

}
func (app *application) Logout(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "localhost", true, true)
	c.Redirect(http.StatusFound, "/")
}
