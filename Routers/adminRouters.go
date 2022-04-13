package Routers

import (
	"123/controllers/admin"
	"123/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRouters(r *gin.Engine) {
	//r.LoadHTMLGlob("templates/**/*")
	adminRouters := r.Group("/admin", middlewares.UserMiddle)
	//adminRouters.Use()
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "我是admin")
		})
		//adminRouters.GET("/user", func(c *gin.Context) {
		//	c.HTML(http.StatusOK,"admin/pull-away-test.html",gin.H{
		//		"user":"我是一个客户",
		//	})
		//})
		adminRouters.GET("/password", admin.UserController{}.Password)
		adminRouters.GET("/age", admin.UserController{}.Age)
	}
}
