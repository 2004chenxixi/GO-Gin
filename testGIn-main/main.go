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

//testGin-html/after-picture/afterPicture.html
//testGin_gorm/gorm.go
//testGin-html/default/num-One.html
//static/css/style.css
//testGin-html/details/registerDetails.html
//testGin-html/resultOne/gain-result.html
//testGin-html/changePassword/changePassword.html
//testGin-html/sign/sign-jump.html
//testGin-html/two-page/page.html
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
	}
	if a == "" || p == "" || a == "NULL " || p == "NULL" {
		ok = false
	} else {
		ok = false
	}
	fmt.Println("请输入你的账号与密码")
	return
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("testGin-html/**/*")

	r.StaticFS("/static-file", http.Dir("./static"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "default/num-One.html", gin.H{})
	})

	r.POST("/User", func(c *gin.Context) {
		if c.PostForm("account") == "" || c.PostForm("password") == "" || c.PostForm("account") == "NULL" || c.PostForm("password") == "NULL" {
			c.String(http.StatusOK, "请输入，还没输入任何东西")
			return
		}
		yes := Query(c.PostForm("account"), c.PostForm("password"))
		if yes == true {
			fmt.Println("测试成功")
			c.HTML(http.StatusOK, "sign/sign-jump.html", gin.H{})

		} else {
			c.String(http.StatusOK, "对不起，你不是该用户")
			fmt.Println("失败")
			return
		}
	})

	r.GET("/gain", func(c *gin.Context) {
		c.HTML(http.StatusOK, "resultOne.html/gain-result.html", gin.H{})
	})
	r.POST("/gain-Picture", func(c *gin.Context) {
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
		fmt.Println("传入成功")
		c.HTML(http.StatusOK, "after-picture/afterPicture.html", gin.H{})
		//c.JSON(http.StatusOK, gin.H{
		//	"gainResult": true,
		//	"dst":        Dst,
		//})
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
		if c.PostForm("account") == WebSlice.Account {
			c.String(http.StatusOK, "该账号已经被注册")
			fmt.Println("该账号已经被注册")
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
		testGin_gorm.DB.Where("account = ?", c.PostForm("account")).Update("password", c.PostForm("password"))
		WebSlice.Password = c.PostForm("newpassword")
		goto delete

	delete:
		testGin_gorm.DB.Where("account = ?", c.PostForm("account")).Delete(&WebSlice)

		if c.PostForm("account") == "" || c.PostForm("password") == "" || c.PostForm("account") == "NULL" || c.PostForm("password") == "NULL" {
			c.String(http.StatusOK, "请输入，还没输入任何东西")
			return
		}
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
