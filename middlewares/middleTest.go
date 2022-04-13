package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func UserMiddle(c *gin.Context) {
	//这边--出现这两个，说明，用户登陆了
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username", "张三")
}
