package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//配置session的中间件
	store := cookie.NewStore([]byte("secret"))   //创建基于cookie的储存引擎，secret是参数用于加密的密钥
	r.Use(sessions.Sessions("mysession", store)) //配置session的中间件，store是存储引擎也可以替换
	r.GET("/", func(c *gin.Context) {

		//	c.SetCookie("username", "张三", 3600, "/", "localhost", false, true)
	})

	r.GET("/c", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("username", "我是session的test")
		session.Save() //设置session时候必须调用
		//username ,_ := c.Cookie("username")
		//c.String(http.StatusOK, "我是第一个的哦"+username)
		//c.String(http.StatusOK, "xiixixixi")
	})

	r.GET("/z", func(c *gin.Context) {

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
