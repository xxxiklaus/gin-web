/* 请求参数 */

package main

import "github.com/gin-gonic/gin"

func TestGet(c *gin.Context) {
	s := c.Query("username")

	pwd := c.DefaultQuery("password", "123")
	c.String(200, "username:%s", "password:%s", s, pwd)
}

func TestPathParam(c *gin.Context) { //获取路径
	s := c.Param("name")
	s2 := c.Param("age")

	c.String(200, "Name:%s,Age:%s", s, s2)

}

func GoSearch(c *gin.Context) {
	c.HTML(200, "query.html", nil)
}

func Search(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	key := c.PostForm("key")
	c.String(200, "page:,key:", page, key)
}

func main() {
	e := gin.Default()
	// e.GET("/testGet", TestGet)
	// e.GET("/testPath/:name/:age", TestPathParam)

	e.LoadHTMLGlob("templates/*")
	e.GET("/goSearch", GoSearch)
	e.POST("/search", Search)

	e.Run(":8888")
}
