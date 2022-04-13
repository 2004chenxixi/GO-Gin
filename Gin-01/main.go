package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "xixiixix")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面，没想到吧")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "我是一个Post请求返回数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "我用来修改数据哦")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "我用来删除数据哦")
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
