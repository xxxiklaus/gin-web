package main

import "github.com/gin-gonic/gin"

func Handler(c *gin.Context) {
	s, err := c.Cookie("username")
	if err != nil {
		s = "klaus"
		c.SetCookie("username", s, 60*60, "/", "localhost", false, true)
	}
	c.String(200, "testcookie")
}

func main() {
	e := gin.Default()

	e.GET("/test_cookie", Handler)
	e.Run()

}
