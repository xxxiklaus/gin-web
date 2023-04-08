package main

import "github.com/gin-gonic/gin"

func Login(context *gin.Context) {
	context.HTML(200, "login.html", nil)
}

func DoLogin(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")
	context.HTML(200, "welcome.html", gin.H{
		"username": username,
		"password": password,
	})
}

func main() {
	e := gin.Default()

	e.LoadHTMLGlob("templates/*") //相对路径

	e.GET("/login", Login)
	e.POST("/login", DoLogin)
	e.Run(":8787")
}
