/*
	Gin实现路由分组

如果你的网站有多个模块:博客、教程、视频、问答,每个模块又有多个路由,这样
就可以进行路由分组,使用的方法是router.Group(分组名称)
*/
package main

import "github.com/gin-gonic/gin"

func F1(c *gin.Context) {
	c.String(200, "blog list")
}
func F2(c *gin.Context) {}
func F3(c *gin.Context) {}
func F4(c *gin.Context) {
	c.String(200, "vdieo list")
}
func F5(c *gin.Context) {}
func F6(c *gin.Context) {}

func main() {
	router := gin.Default()

	//博客 访问: http://localhost:8080/blog/list
	v1 := router.Group("/blog")
	{
		v1.GET("/list", F1)
		v1.POST("/post", F2)
		v1.POST("/add", F3)
	}
	//视频 访问: http://localhost:8080/video/list
	v2 := router.Group("/video")
	{
		v2.GET("/list", F4)
		v2.POST("/post", F5)
		v2.POST("/add", F6)

	}
	router.Run()

}
