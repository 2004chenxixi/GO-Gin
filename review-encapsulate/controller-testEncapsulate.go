package controllerTestEncapsulate

import (
	tool "123/models"
	"fmt"
	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

type Controller struct {
}

func (con Controller) EncapsulateGet(c *gin.Context) {
	//设置session

	c.HTML(http.StatusOK, "admin/user.html", gin.H{})
}

func (con Controller) EncapsulatePost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	age := c.DefaultPostForm("age", "16未成年")
	//按日期上传和储存文件，一共5步骤
	//1.获取文件
	file, err := c.FormFile("face")
	//2.判断后缀名
	extName := path.Ext(file.Filename) //path.Ext->可以获取文件后缀名--.jpg  .png  .gif   .jpeg(这些自定义)
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
	//3.创建-当天的图片，放在当天的目录下
	day := tool.GetDay()
	dir := "./static/upload" + day

	os.MkdirAll(dir, 0777) //就是用os.MkdirAll来创建目录,这里后面的0777为权限，要注意⚠️
	if err != nil {
		fmt.Println("创建目录有误")
		c.String(http.StatusOK, "创建目录失败，请从新尝试")
		return
	}
	//4.对应目录下生成的文件
	//文件名应该是:timeUnix + extName,但是timeUnix为int64类型，要先转化strconv.FormatInt()这一类可以接收int64类型转化为string
	fileName := strconv.FormatInt(tool.GetUnix(), 10) + extName
	//5.上传
	dst := path.Join(dir, fileName) //这里的意思是，将filename文件传到dir中，dir是上面自己创建到目录
	fmt.Println("dst", dst)
	c.SaveUploadedFile(file, dst) //5.上传

	c.JSON(http.StatusOK, gin.H{
		"gainResult": true,
		"username":   username,
		"password":   password,
		"age":        age,
		"dst":        dst,
	})

}
