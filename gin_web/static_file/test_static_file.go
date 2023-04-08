/* Gin访问静态文件集成BootStrap框架
下载BootStrap https://getbootstrap.com/
添加bootstrap css和js文件
创建一个assets文件夹,将css和js文件添加到该文件夹中 */

package main

import "github.com/gin-gonic/gin"

func Go(e *gin.Context) {
	e.HTML(200, "static.html", nil)
}

func main() {
	e := gin.Default()
	e.Static("/assets", "./assets")
	e.LoadHTMLGlob("templates/*")
	e.GET("/go", Go)
	e.Run()

}
