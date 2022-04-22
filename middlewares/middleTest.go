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
	c.Set("username", "张三") //在中间间中设置值
	//再使用c.Get(上面的"username")就可以在需要的(终端)输入出来
	//但是要是想在(网页)输出出来，需要类型断言，把它变成string类型的

	//这里的细节是要把cCp拿出来copy一下，，，才可以使用
	cCp := c.Copy()
	//使用goroutine
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done  in path" + cCp.Request.URL.Path)
	}()
}
