package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Println(str01 string, str02 string) string { //自定义函数
	return str01 + str02
}

//定义一个GET-POST的结构体
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password"  form:"password"`
	Age      int    `json:"age" form:"age"`
}

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content"xml:"content"`
}

func main() {
	//创建一个路由引擎
	r := gin.Default()

	//自定义模版函数---把这个函数放在加载模版前
	r.SetFuncMap(template.FuncMap{
		//这个地方用来接受(固定-->r.SetFuncMap(template.FuncMap{ )，
		// 自己写的方法，要放在r.LoadHTMLGlob("templates/**/*")之前
		"Println": Println,
	})

	//加载模版---放在配置路由前
	//下面--用来引用html
	r.LoadHTMLGlob("templates/**/*")

	//--下面--配置静态web服务  第一参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")

	//GET请求传值
	r.GET("/art", func(c *gin.Context) {
		//可以从外部传值--？后+需要的东西
		//c.DefaultQuery这个可以自己写
		id := c.DefaultQuery("id", "1")

		c.JSON(http.StatusOK, gin.H{
			"msg": "新闻页面",
			"id":  id,
		})
	})
	//POST演示*******
	r.GET("/user2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/user.html", gin.H{})
	})
	//获取表单post过来的数据
	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "16未成年")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	//获取GET--传递的数据绑定到结构体
	r.GET("/user", func(c *gin.Context) {
		user := &UserInfo{}
		// c.ShouldBind表示获取数据
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": "err.Error",
			})
		}
	})
	//获取POST--接受传递的数据绑定到结构体
	//这里的/doAddUser2就是钥匙，可以把html输入的内容导过来
	r.POST("/doAddUser2", func(c *gin.Context) {
		postUser := &UserInfo{}
		if err := c.ShouldBind(&postUser); err == nil {
			c.JSON(http.StatusOK, postUser)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": "err.Error",
			})
		}
	})

	r.POST("/xml", func(c *gin.Context) {
		article := &Article{}
		xmlSlicedate, _ := c.GetRawData() //从c.Request.Body读取数据
		//将字符串变成结构体，用xml.Unmarshal，第一个写字符串，第二个写想变的结构体

		fmt.Println(xmlSlicedate)
		if err := xml.Unmarshal(xmlSlicedate, &article); err == nil {
			c.JSON(http.StatusOK, xmlSlicedate)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": "Error",
			})
		}
	})

	//动态路由
	r.GET("list/:cid", func(context *gin.Context) {
		cid := context.Param("cid")
		context.String(http.StatusOK, "%v", cid)
	})
	r.Run(":9090")
}
