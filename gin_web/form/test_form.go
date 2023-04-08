/* Gin表单处理 */

package main

import "github.com/gin-gonic/gin"

func GoRegister(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	hobby := c.PostFormArray("hobby")
	gender := c.PostForm("gender")
	city := c.PostForm("city")

	c.String(200, "Username:%s,Password:%s,hobby:%s,gender:%s,city:%s",
		username, password, hobby, gender, city)
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.POST("/register", Register)
	e.GET("/testForm", GoRegister)
	e.Run(":8888")
}
