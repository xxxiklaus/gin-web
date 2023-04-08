/* Gin 使用中间件
/在请求中间起到拦截作用的处理函数
/ Gin默认中间件
/ 如果使用的是Gin.Default()实例化gin引擎,默认有两个中间件,Logger和Recovery,分别用来处理日志和处理错误
/如果使用gin.New()需要重新添加
r := gin.New()

/全局中间件
/Logger 中间件将日志写入gin.DefaultWriter,即使你将GIN_MODE设置为release.
/By default gin.DefaultWriter = os.Stdout

r.Use(gin.Logger())
/Recovery中间件会recover任何panic.如果有panic的话,会写入500.
r.Use(gin.Recovery()) */

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMW(c *gin.Context) {
	c.String(200, "hello Golang!", nil)
}

func MyMiddleware(c *gin.Context) {
	fmt.Println("我的第一个中间件")
}

func MyMiddleware2(c *gin.Context) {
	fmt.Println("我的第二个中间件")
}

func main() {

	e := gin.Default()
	/* 	func Default() *Engine {
		debugPrintWARNINGDefault()
		engine := New()
		engine.Use(Logger(), Recovery()) 默认有Logger和Recovery两个中间件
		return engine
	} */
	e.Use(MyMiddleware, MyMiddleware2)

	e.GET("testmw", TestMW)

	e.Run()

}
