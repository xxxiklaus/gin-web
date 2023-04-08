/* Gin是一个golang微框架,基于httprouter,封装比较优雅,API友好
源码注释比较明确,具有快速灵活,容错方便等特点.
特征
速度快:基于基数树的路由,内存占用少.没有反射.可预测API性能
中间件支持:传入的HTTP请求可以由中间件链和最终操作处理.例如:Logger、
Authorization、GZIP最后在DB中的一条消息
Crash-free: Gin可以捕获HTTP请求期间发生的panic并恢复它.保持服务器始终可用
路由分组: 更好地组织你的路线.需要授权与不需要授权,不同的API版本.此外
组可以无限嵌套,而不会降低性能
错误管理: Gin提供了一种方便的方法来收集HTTP请求期间发生的所有错误.最终
中间件可以将它们写入日志文件数据库并通过网络发送它们
内置渲染: Gin为JSON、XML和HTML渲染提供了一个易于使用的API
可扩展: 创建一个新的中间件非常简单,只需要查看示例代码. */

package main

import "github.com/gin-gonic/gin"

func Hello(context *gin.Context) { //路由函数
	context.String(200, "hello,%s, golang!")
}

func main() {
	router := gin.Default() //创建一个默认的路由

	router.GET("/hello", Hello) //绑定路由规则和路由函数,访问/hello的路由,将由对应的函数取处理
	router.Run(":8787")         //启动监听,gin会把服务运行在本机的 0.0.0.0(可以是本机的任何一个ip)默认8080

}

/* package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
} */
