package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Middle(c *gin.Context) {
	fmt.Println("我是第一个中间件")
	start := time.Now().UnixNano()

	c.Next() //调用该请求的剩余处理---不会直接到c.Next会先回到main中，执行完main，然后再执行c.Next后面的
	//c.Abort--只会执行方法内的，不会回到main中，与c.Next是相反的存在

	fmt.Println("我是第二个中间件")
	end := time.Now().UnixNano()

	fmt.Println("start-end=result", start-end)
}

func MiddleOne(c *gin.Context) {
	fmt.Println("我是第一个中间件--one")

	c.Next() //调用该请求的剩余处理---不会直接到c.Next会先回到main中，执行完main，然后再执行c.Next后面的
	//c.Abort--只会执行方法内的，不会回到main中，与c.Next是相反的存在

	fmt.Println("我是第二个中间件--one")
}
func MiddleTwo(c *gin.Context) {
	fmt.Println("我是第一个中间件-two")

	c.Next() //调用该请求的剩余处理---不会直接到c.Next会先回到main中，执行完main，然后再执行c.Next后面的
	//c.Abort--只会执行方法内的，不会回到main中，与c.Next是相反的存在

	fmt.Println("我是第二个中间件-two")

}

func main() {
	r := gin.Default()

	r.Use(MiddleOne, MiddleTwo)
	r.GET("/", func(c *gin.Context) {
		fmt.Println("one--gin首页")
		c.String(200, "one--gin首页")
	}, func(c *gin.Context) {
		fmt.Println("two--gin首页")
		c.String(200, "two---gin首页")
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

//c.String(200,"gin首页")
