package Routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/API")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "我是API接口")
		})
		apiRouters.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK, "我是API接口--user")
		})
		apiRouters.GET("/password", func(c *gin.Context) {
			c.String(http.StatusOK, "我是API接口--password")
		})
		apiRouters.GET("/age", func(c *gin.Context) {
			c.String(http.StatusOK, "我是API接口--age")
		})
	}
}
