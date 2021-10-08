package main

import (
	"blog/pkg"
	"github.com/gin-gonic/gin"
)

func(app *application) routes() *gin.Engine  {
	route := gin.Default()
	route.LoadHTMLGlob("./ui/template/*")

		route.GET("/",app.Home)
		route.GET("/signpage",app.SignUpPage)
		route.POST("/signup",app.Signup)
		route.GET("/loginpage",app.LoginPage)
		route.POST("/login",app.Login)



	subRoute := route.Group("/user")
	subRoute.Use(pkg.CheckLogin())
	{
		subRoute.POST("/addpost",app.AddPost)
		subRoute.GET("/addpostpage",app.AddPostPage)
	}

	return route
}