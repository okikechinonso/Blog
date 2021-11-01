package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	route := gin.Default()
	route.Static("/assets", "./ui/assets")
	route.LoadHTMLGlob("./ui/template/*")

	route.GET("/", app.Home)
	route.GET("/signup", app.SignUpPage)
	route.POST("/signup", app.Signup)
	route.GET("/login", app.LoginPage)
	route.POST("/login", app.Login)
	route.GET("/logout", app.Logout)
	route.GET("/blogHome", app.blogHome)
	route.POST("/addpost", app.AddPost)
	route.GET("/addpostpage", app.AddPostPage)
	route.GET("/edit", app.Edit)
	route.GET("/view", app.View)

	return route
}

//subRoute := route.Group("/")
////subRoute.Use(pkg.CheckLogin())
//{
//
//}
