package main

import (
	tools "123/models"
	"123/testGin_gorm"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
)

func Query(a string, p string) (ok bool) {
	WebSlice := testGin_gorm.Web{}
	//testGin_gorm.DB.Find(&WebSlice)
	testGin_gorm.DB.Where("account = ?", a).Find(&WebSlice)
	if WebSlice.Account == a && WebSlice.Password == p {
		ok = true
	} else if a == "" && p == "" {
		ok = false
	} else {
		ok = false
	}
	return
}

func exchangePassword(a string, p string) (ok bool) {
	WebSlice := testGin_gorm.Web{}
	testGin_gorm.DB.Where("account = ?", a).Find(&WebSlice)
	if WebSlice.Account == a && WebSlice.Password == p {
		ok = true
	} else if a == "" || p == "" {
		ok = false
	} else {
		ok = false
	}
	return
}

func create(a string, p string) (ok bool) {

	return
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("testGin-html/**/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "default/num-One.html", gin.H{})
		//WebSlice := testGin_gorm.Web{} //这里用到了gorm，gorm.Bank这个是已经定义好的，结构体在gorm中
		//testGin_gorm.DB.Find(&WebSlice)
		//fmt.Println("看看答案效果=", WebSlice)
	})

	r.POST("/User", func(c *gin.Context) {
		yes := Query(c.PostForm("account"), c.PostForm("password"))
		if yes == true {
			fmt.Println("成功")
			goto come
		} else {
			c.String(http.StatusOK, "对不起，你不是该用户")
			fmt.Println("失败")
			return
		}
	come:
		password := c.PostForm("password")
		account := c.PostForm("account")
		file, err := c.FormFile("face")
		extName := path.Ext(file.Filename)
		allExtMap := map[string]bool{
			".jpg":  true,
			".png":  true,
			".gif":  true,
			".jpeg": true,
		}
		_, ok := allExtMap[extName]
		if ok != true {
			fmt.Println("输入的文件后缀有误")
			c.String(http.StatusOK, "你的输入的照片后缀有误，请从新尝试")
			return
		}
		dir := "./static/upload" + tools.DAY()
		os.MkdirAll(dir, 0777) //就是用os.MkdirAll来创建目录,这里后面的0777为权限，要注意
		if err != nil {
			fmt.Println("创建目录有误")
			c.String(http.StatusOK, "创建目录失败，请从新尝试")
			return
		}
		fileName := tools.Day() + extName
		Dst := path.Join(dir, fileName)
		c.SaveUploadedFile(file, Dst)
		c.JSON(http.StatusOK, gin.H{
			"gainResult": true,
			"password":   password,
			"account":    account,
			"dst":        Dst,
		})
	})

	r.GET("/exchange", func(c *gin.Context) {
		c.HTML(200, "two-page/page.html", gin.H{
			"我是一个跳转到注册页面": "跳转注册页面",
		})
	})
	r.POST("/link", func(c *gin.Context) {
		WebSlice := testGin_gorm.Web{
			Account:  c.PostForm("account"),
			Password: c.PostForm("password"),
		}
		if c.PostForm("account") == "" || c.PostForm("password") == "" || c.PostForm("account") == "NULL" || c.PostForm("password") == "NULL" {
			c.String(http.StatusOK, "请输入，还没输入任何东西")
			return
		}
		testGin_gorm.DB.Create(&WebSlice)
		c.String(http.StatusOK, "恭喜，注册成功")
	})
	r.GET("/details", func(c *gin.Context) {
		c.HTML(200, "details/registerDetails", gin.H{
			"我是一个注册细节": "细节",
		})
	})

	r.GET("/changePassword", func(c *gin.Context) {
		c.HTML(200, "changePassword/changePassword.html", gin.H{
			"我是一个修改密码": "修改密码",
		}) //显示成功，但是不可以在--数据库该密码
	})
	r.POST("/password", func(c *gin.Context) {
		change := exchangePassword(c.PostForm("account"), c.PostForm("password"))
		WebSlice := testGin_gorm.Web{Account: c.PostForm("account")}
		testGin_gorm.DB.Where("account = ?", c.PostForm("account")).Find(&WebSlice)
		testGin_gorm.DB.Delete(c.PostForm("account"))
		WebSlice.Password = c.PostForm("newpassword")

		if change == true {
			c.String(http.StatusOK, "修改成功")
			testGin_gorm.DB.Save(&WebSlice)
			//testGin_gorm.DB.Save(&WebSlice)
		} else {
			c.String(http.StatusOK, "修改失败")
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

//testGin-html/default/num-One.html
//models/tools.go
//testGin_gorm/gorm.go
//testGin_gorm/webGorm.go

//func init() {
//	for i := 0; i < 3; i++ {
//		WebSlice := []testGin_gorm.Web{}
//		testGin_gorm.DB.Find(&WebSlice)
//		fmt.Println(WebSlice)
//	}
//}

//r.POST("/password", func(c *gin.Context) {
//	change := exchangePassword(c.PostForm("account"), c.PostForm("password"))
//	WebSlice := testGin_gorm.Web{Account: c.PostForm("account")}
//	//testGin_gorm.DB.Update("password", WebSlice)
//	testGin_gorm.DB.Find(&WebSlice)
//	WebSlice.Password = c.PostForm("newpassword")
//	if change == true {
//		testGin_gorm.DB.Save(&WebSlice)
//		c.String(http.StatusOK, "修改成功")
//	} else {
//		c.String(http.StatusOK, "修改失败")
//	}
//})
//
