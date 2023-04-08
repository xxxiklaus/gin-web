package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type User struct { //数据绑定
	Username string   `form:"username"`
	Password string   `form:"password"`
	Hobby    []string `form:"hobby"`
	Gender   string   `form:"gender"`
	City     string   `form:"city"`
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "form.html", nil)
}

func Register(c *gin.Context) {
	var user User
	c.ShouldBind(&user)
	c.String(200, "form data:%s", user)
}

type User2 struct { //查询数据参数
	Username string `form:"username"`
	Password string `form:"password"`
}

func TestGetBind(c *gin.Context) {
	var user User2
	err := c.ShouldBind(&user)
	if err != nil {
		log.Fatal(err)
	}
	c.String(200, "User:%s", user)
}

type User3 struct { //查询路径参数
	Username string `uri:"username"`
	Password string `uri:"password"`
}

func TestUriBind(c *gin.Context) {
	var user User3
	err := c.ShouldBindUri(&user)
	if err != nil {
		log.Fatal(err)
	}
	c.String(200, "User:%s", user)
}

func main() {
	e := gin.Default()

	e.LoadHTMLGlob("templates/*")
	e.GET("/register", GoRegister)
	e.POST("/register", Register)
	//http://localhost:8888/testGetBind?username=klaus&password=123
	e.GET("/testGetBind", TestGetBind)
	//http://localhost:8888/testUriBind/klaus/123
	e.GET("/testUriBind/:username/:password", TestUriBind)

	e.Run(":8888")

}
