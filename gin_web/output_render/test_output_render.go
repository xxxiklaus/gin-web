/*
	Gin输出渲染

Gin支持很多种输出渲染,可以是简单的字符串、JSON、XML、HTML、ProtoBuf.
c.JSON(200,nil)
c.XML(200,nil)
c.HTML(200,"",nil)
c.String(200,"")
c.ProtoBuf(200,nil)
*/
package main

import "github.com/gin-gonic/gin"

func TestJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "Github",
		"site": "wwww.github.com",
	})
}

func TestXML(c *gin.Context) {
	c.XML(200, gin.H{
		"name": "Github",
		"site": "wwww.github.com",
	})
}

func TestHTML(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func TestString(c *gin.Context) {
	c.String(200, "github,xxxiklaus README")
}

func main() {
	e := gin.Default()

	e.GET("/test_json", TestJSON)
	e.GET("/test_xml", TestXML)
	e.LoadHTMLGlob("templates/*")
	e.GET("/test_html", TestHTML)
	e.GET("/test_string", TestString)

	e.Run()

}
