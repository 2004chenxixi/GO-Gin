## gin框架-----数据库与Go连接

## 主要思路

```go-skill
1--方法可以查GIn的官网,只要改密码和连接的表就可以,一般情况会写(一个连接and一个结构体)

2--连接完,可以对,数据库,增删改查-----具体看testGIn-main(这个是自己写的,可以查看)

3--但是要是想网页输入东西,到数据库-----需要GET--再POST接受

4--可以先引入--r.LoadHTMLGlob("testGin-html/**/*")---才可以使用html--POST接受,是写再html里面的----那个具体内容,根据自定义的目录改
** 注释4--如果你照抄,4的路径-----要把主要目录设置成testGin-html--然后在他"下面需要一个，就新建一个目录" 
such as:
testGin-html(项目下面的主目录)
|
你需要的html的小目录+你的html
|
这个还是你需要的html放的小目录+你的html
|
这里也是小目录+你要的html
//总结，他们都是在总目录下的若干个小目录

5--r.StaticFS("/static-file", http.Dir("./static"))
如果要引入图片或者css这些东西,就要引入5---前面的"/static-file"可以自己定义,他只是后面的映射,后面的
"./static"是放img,css这些东西
such as:
static(主目录)
|
css(小目录)+xxx.css(file)
|
images(小目录)+xxx.文件后缀
！！！！！！！！！！！！！！！！！！！！！
**********************************
重要总结==想用html，    引入r.LoadHTMLGlob
重要总结==想要图片，css，引入r.StaticFS
路径要对，不然没用
***********************************
```

## 数据库和GO连接

```go-
数据库和GO连接
｜｜
1-
package testGin_gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:ckj20040319@tcp(localhost:3306)/webTest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
}

2-写一个结构体和方法，连接数据库
这里的结构体名---是"数据库连接Go时的那个表的名"同名即可
这里的结构体内容---是"数据库连接Go时的那个表里面的基础条件----同样的"基础结构"就可以
｜｜
package testGin_gorm

type Web struct {
	Id       string `gorm:"column:Id;Primary_key""`
	Account  string
	Password string
}

func (Web) TableName() string {
	return "web"
}

｜｜
注意==这里我是吧他们放在testGin_gorm同一个目录下面就可以
目录结构如下
｜｜
testGin_gorm(主目录)
｜
gorm.go(file)
|
webGorm.fo(file)
```

## 主包的代码

```go-
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
			c.HTML(http.StatusOK,"sign/sign-jump.html",gin.H{})

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
			c.HTML(http.StatusOK,"after-picture/afterPicture.html",gin.H{})
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
		if c.PostForm("account") == WebSlice.Account{
			c.String(http.StatusOK,"该账号已经被注册")
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

主要路径
//testGin-html/after-picture/afterPicture.html
//testGin_gorm/gorm.go
//testGin-html/default/num-One.html
//static/css/style.css
//testGin-html/details/registerDetails.html
//testGin-html/resultOne/gain-result.html
//testGin-html/changePassword/changePassword.html
//testGin-html/sign/sign-jump.html
//testGin-html/two-page/page.html

路径
｜｜
testGin-main(目录)
｜
main.go(文件)

```

## html的代码路径

##### 下面会一次将对应的代码发出来

```go-
先说目录
testGin-html(主要目录)
｜
after-picture(小目录)+afterPicture.html(file)
｜
changePassword(小目录)+changePassword.html(file)
|
default(小目录)+num-One.html(file)
|
details(小目录)+registerDetails.html(file)
|
resultOne(小目录)+gain-result.html(file)
|
sign(小目录)+sign-jump.html(file)
|
two-page(小目录)+page.html(file)
```

#### afterPicture.html的代码

```go-
{{define "after-picture/afterPicture.html"}}
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Title</title>
    </head>
    <body>
    <h1>传入成功</h1>
    <a href="http://localhost:8080/exchange" target="_blank">下一页</a>
    </body>
    </html>
{{end}}
```

#### changePassword.html的代码

```go-
{{define "changePassword/changePassword.html"}}
    <!DOCTYPE html>
    <html lang="en">
    <link rel="stylesheet" href="/static-file/css/style.css">
    <div class="password">
    <head>
        <meta charset="UTF-8">
        <title>Title</title>
    </head>
    <body>
    <form action="/password" method="post" enctype="multipart/form-data">
        <div class="care-change">
        <p>友情提示:请不要随便修改密码，如果一定要修改，请将账号和密码一次输入正确，如果<em>显示修改错误</em>，系统将直接注销你的账号,残忍ing</p>
        </div>
        你的账号<input type="text" name="account"/><br><br>
        旧:密码<input type="text" name="password"/><br><br>
        新:密码 <input type="text" name="newpassword"/><br><br>
        <input type="submit" value="修改">
    </form>
    </body>
    </div>
    </html>
{{end}}
```

#### num-One.html的代码

```go-
{{ define  "default/num-One.html"}}
    <!DOCTYPE html>
    <html xmlns="http://www.w3.org/1999/html">
    <link rel="stylesheet" href="/static-file/css/style.css">
    <head>
        <title>HTML</title>
    </head>

    <body>
    <div class="border">
        <div class="title">
            <h1>MYDX</h1>
            <h2>__________________________</h2>
            <p><em>如果没有账号，请注册</em>
                <a href="http://localhost:8080/exchange" target="_blank">注册</a></p>
        </div>

        <br>
        <div class="middle">

            <form action="/User" method="post" enctype="multipart/form-data">
                账号<input type="text" name="account"/><br><br>
                <p>密码<input type="password" name="password"/>
                    <a href="http://localhost:8080/changePassword" target="_blank">修改密码</a></p>
                <input type="submit" value="登入测试">
            </form>
        </div>
    </div>

    </body>
    <hr>
    <br>

    <br>

    {{/*    <img src="/static-file/images/view.jpg" width="100p" alt="图片输入错误">*/}}
    <div class="end">

        <blockquote>如果有任何技术问题，自行解决，本网站不支持任何解决方案</blockquote>


        <blockquote><abbr title="自己查百度">修复BUG</abbr>,如果出现BUG</blockquote>
    </div>

    </html>



    <br>
{{end}}
```

#### registerDetails.html的代码

```go-
{{define  "details/registerDetails"}}
    <!DOCTYPE html>
    <html lang="en">
    <link rel="stylesheet" href="/static-file/css/style.css">
    <div class="care">
    <head>
        <h1>欢迎来到注册细节</h1>
    </head>
    <body>
    <p>我们注册分为"两步骤"</p>
    <ol>
        <li>设置一个账号名称</li>
        <li>设置密码</li>
    </ol>
    </body>
    </div>
    </html>
{{end}}
```

#### gain-result.html的代码

```go-
{{define "resultOne.html/gain-result.html"}}
    <!DOCTYPE html>
    <html xmlns="http://www.w3.org/1999/html">
    <link rel="stylesheet" href="/static-file/css/style.css">
    <div class="result">
    <head>
        <title>HTML</title>
    </head>
    <body>
    <div class="picture">
        <p>注意这边，你一定要先传入一张照片，才能继续哦</p>
        <p>骗照片我们是认真的</p>
        <p>请按照流程走</p>
    </div>
    <h1>欢迎进入--请先输入一张美美的照片吧</h1>
    <form action="/gain-Picture" method="post" enctype="multipart/form-data">
        照片<input type="file" name="face"/><br><br>
        <input type="submit" value="提交">
    </form>

    </body>
       </div>
    </html>
{{end}}
```

#### sign-jump.html的代码

```go-
{{define "sign/sign-jump.html"}}
    <!DOCTYPE html>
    <html lang="en">
    <link rel="stylesheet" href="/static-file/css/style.css">
    <div class="sign">
        <head>
            <meta charset="UTF-8">
            <title>Title</title>
        </head>
        <body>
        <h3>恭喜你通过--登入测试</h3>
        <h3>请点击"登入"直接进入</h3><a href="http://localhost:8080/gain" target="_blank">登入</a>

        </body>
    </div>
    </html>
{{end}}
```

#### page.html的代码

```go-
{{define "two-page/page.html"}}
    <!DOCTYPE >
    <html lang="en">
    <link rel="stylesheet" href="/static-file/css/style.css">
    <div class="register">
        <head>
            <meta charset="UTF-8">
            <title>Title</title>
        </head>
        <h1>欢迎来到注册页面</h1>
        <body>
        <p>友情提示:建议先看注册细节</p>
        <p><em>注册细节</em><a href="http://localhost:8080/details" target="_blank">注册细节</a></p>
        <br>
        <form action="/link" method="post" enctype="multipart/form-data">
            创建:账号<input type="text" name="account"/><br><br>
            创建:密码<input type="text" name="password"/><br><br>
            <input type="submit" value="注册">
        </form>
        </body>
    </div>
    </html>
{{end}}
```

## static的路径

##### 代码在下面

```go-
static(主目录)
|
css(小目录)+style.css(下面有这个的代码)
|
images(小目录)+xxx.图片后缀
|
js(小目录)+文件
|
upload(小目录)---如果你有看主代码,就知道--这个会更具每次上传的时间自己创建一个目录
```

#### style.css的代码

#### 用来和html合作美化页面

```go-
.title {
    width: 30%;
    margin: auto;
    color: dodgerblue;
}

.middle {
    width: 30%;
    margin: auto;
}

.end {
    width: 30%;
    margin: auto;
    border: 10px dotted #ba7235;
    background: #ba7235;
}

.border {
    border: 10px dotted #51862c;
    background: #ba7235;
}

.care {
    width: 30%;
    margin: auto;

}
.result{
    width: 30%;
    margin: auto;
}
.password{
    width: 30%;
    margin: auto;
}
.sign{
    width: 30%;
    margin: auto;
}
.register{
    width: 30%;
    margin: auto;
}
.care-change{
    background: red;
    color: whitesmoke;
}
.picture{
    background: red;
    color: whitesmoke;
}
```

## 自己写的辅助工具

```go-
路径
｜｜
models(主要目录)
｜
tools.go(文件)
```

#### tools.go的代码

```go-
package tools

import (
	"time"
)

func GetDay() string {
	test := "1650422138"
	return time.Now().Format(test)
}
func GetUnix() int64 {
	return time.Now().Unix()
}

func Day() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	return timeStr
}
func DAY() string {
	timeStr := time.Now().Format("2006-01-02")
	return timeStr
}

```















