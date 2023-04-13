# gin-web
gin_web basic

一些基础的实现和应用  
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{ //模拟一些私人数据
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@xample.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@uapa.com", "phone": "523433"},
}

func Handler(c *gin.Context) {
	//获取用户,它是由BasicAuth中间件设置
	user := c.MustGet(gin.AuthUserKey).(string)
	fmt.Println(user)
	if secret, ok := secrets[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :( "})
	}
}

func main() {
	e := gin.Default()

	e1 := e.Group("/admin", gin.BasicAuth(gin.Accounts{ //这些用户数据可以从数据库gorm查询
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	//http://localhost:8080/admin/secrets
	e1.GET("/secrets", Handler)
	e.Run()
}
```
