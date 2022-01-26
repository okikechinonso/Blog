package main

import (
	"blog/models"

	"github.com/gin-gonic/gin"
)

type Application struct {
	user models.IUsers
	blog models.Iblog
	
}

func (app *Application) routes() *gin.Engine {
	route := gin.Default()
	route.LoadHTMLGlob("./ui/template/*")

	route.GET("/", app.Home)
	route.GET("/signpage", app.SignUpPage)
	route.POST("/signup", app.Signup)
	route.GET("/loginpage", app.LoginPage)
	route.POST("/login", app.Login)
	route.GET("/logout", app.Logout)
	route.GET("/blogHome", app.blogHome)

	subRoute := route.Group("/user")
	//subRoute.Use(pkg.CheckLogin())
	{
		subRoute.POST("/addpost", app.AddPost)
		subRoute.GET("/addpostpage", app.AddPostPage)
	}

	return route
}
