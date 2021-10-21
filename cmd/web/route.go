package main

import (
	"github.com/gin-gonic/gin"
)

func(app *application) routes() *gin.Engine  {
	route := gin.Default()
	route.Static("/assets", "./ui/assets")
	route.LoadHTMLGlob("./ui/template/*")

		route.GET("/",app.Home)
		route.GET("/signpage",app.SignUpPage)
		route.POST("/signup",app.Signup)
		route.GET("/loginpage",app.LoginPage)
		route.POST("/login",app.Login)
		route.GET("/logout",app.Logout)
		route.GET("/blogHome",app.blogHome)



	subRoute := route.Group("/user")
	//subRoute.Use(pkg.CheckLogin())
	{
		subRoute.POST("/addpost",app.AddPost)
		subRoute.GET("/addpostpage",app.AddPostPage)
	}

	return route
}