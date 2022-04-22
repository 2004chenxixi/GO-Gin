package admin

import (
	"123/baseContrllers"
	"fmt"
	"net/http"
	"path"

	//"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	baseContrllers.BaseControllers
}

func (con UserController) Password(c *gin.Context) {
	//c.String(http.StatusOK, "我是admin--password--我是func分装的")

	con.Success(c)
}
func (con UserController) Age(c *gin.Context) {
	//c.String(http.StatusOK, "我是admin--age-------我是func分装的")
	//con.Fail(c)
	c.HTML(http.StatusOK, "file-upload/file-upload.html", gin.H{})
}
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username") //这个用来接受POST的username
	//3---------相同名字的多个文件上传-----的方法---c.MultipartForm
	form, _ := c.MultipartForm()
	files := form.File["face[]"]
	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename) //这里的./static/upload是到时候存放收入进来的文件的地方
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"gainResult": true,
		"username":   username,
	})
	//1---------下面是---单个文件上传
	//username := c.PostForm("username") //这个用来接受POST的username
	//file, err := c.FormFile("face")    //这个用来接受POST传进来的文件-file类型的
	////file.Filename--->来获取文件名称，这样对下面的dst有用
	//dst := path.Join("./static/upload", file.Filename) //这里的./static/upload是到时候存放收入进来的文件的地方
	//if err != nil {
	//	fmt.Println("err")
	//} else {
	//	c.SaveUploadedFile(file, dst)
	//}
	//c.String(200, "执行上传")
	//c.JSON(http.StatusOK, gin.H{
	//	"gainResult": true,
	//	"username":   username,
	//	"dst":        dst,
	//})
	//上传单个文件的上传方法
	//1、传入文件file, err := c.FormFile("face") 这里要if来对付err所以把【2】放到if外，把【3】放到if正确的输出
	//这个用来接受POST传进来的文件-file类型的
	//2、dst := path.Join("./static/upload", file.Filename) //这里的./static/upload是到时候存放收入进来的文件的地方
	//3、c.SaveUploadedFile(file, dst) ******----->3用来是接受1与2
	//4、c.JSON(http.StatusOK, gin.H{--------->把他们变成字符串打印出来
}

//2--------多个名字，多文件---传输----这个和单文件差不多
func (con UserController) Gender(c *gin.Context) {
	//c.String(http.StatusOK, "我是admin--age-------我是func分装的")
	//con.Fail(c)
	c.HTML(http.StatusOK, "file-upload/file-upload2.html", gin.H{})
}
func (con UserController) DoGender(c *gin.Context) {
	//c.String(200,"执行修改")
	username := c.PostForm("username") //这个用来接受POST的username

	file1, err := c.FormFile("face1") //这个用来接受POST传进来的文件-file类型的
	//file.Filename--->来获取文件名称，这样对下面的dst有用
	dst1 := path.Join("./static/upload", file1.Filename) //这里的./static/upload是到时候存放收入进来的文件的地方
	if err != nil {
		fmt.Println("err")
	} else {
		c.SaveUploadedFile(file1, dst1)
	}

	file2, err := c.FormFile("face2") //这个用来接受POST传进来的文件-file类型的
	//file.Filename--->来获取文件名称，这样对下面的dst有用
	dst2 := path.Join("./static/upload", file2.Filename) //这里的./static/upload是到时候存放收入进来的文件的地方
	if err != nil {
		fmt.Println("err")
	} else {
		c.SaveUploadedFile(file2, dst2)
	}
	//c.String(200, "执行上传")
	c.JSON(http.StatusOK, gin.H{
		"gainResult": true,
		"username":   username,
		"dst1":       dst1,
		"dst2":       dst2,
	})
}
