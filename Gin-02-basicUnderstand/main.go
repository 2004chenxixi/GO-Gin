package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type article struct {
	Title string `json:"title"` //直接在打字模式下，点esc下面那个，打入json就自动跳出
	write string `json:"write"`
}

func main() {

	r := gin.Default()
	//配置模版文件
	r.LoadHTMLGlob("templates/*") //这里的(templates)是与你自己创建的文件对应的

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "没想到吧")
	})

	r.GET("/JSON01", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{ //这个map[string]interface{}--是空的接口，可以打random值
			"我是老六，没想到吧---JSON":   true,
			"我也是老六，想到吧-----JSON": "不服吗",
		})
	})

	r.GET("/JSON02", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{ //这里的gin.H与上面的接口一样用
			"我是老六，没想到吧---JSON02":   false,
			"我也是老六，想到吧-----JSON02": "我佩服了",
		})
	})

	r.GET("/JSON03", func(context *gin.Context) {

		gain := &article{
			Title: "罪与法",
			write: "主人公～罗佳",
		}
		context.JSON(http.StatusOK, gain) //这里的gain是是上面写好的结构体，gain给他附给他一个值，然后只要输出gain就可以
	})

	//JSONP
	r.GET("/JSONP", func(context *gin.Context) {
		againGain := &article{
			Title: "安徒生童话故事",
			write: "主人公～小红毛",
		}
		context.JSONP(http.StatusOK, againGain)
		//这里解释一下JSONP，他本质与JOSON一样，但是她可以在基础上附加东西
		//such-as----->http://localhost:8000/JSONP?callback=陈嘻嘻
		//就是加入?callback=xxx
		//输出结果为---->"title":"安徒生童话故事"
		//变成----->陈嘻嘻嘻({"title":"安徒生童话故事"})
	})

	//XML
	r.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"success":    true,
			"我不知道什么是XML": "xml:我自己也不知道",
		})
	})

	//HTML----这里是渲染模块
	//渲染文件+配置文件
	//以下是配置文件，一定要先配置文件
	//当然在配置文件之前要先写一个可以配置的环境，比如templates里面的-xxx.html这样的环境
	//r.LoadHTMLGlob("templates/*")  //这里的(templates)是与你自己创建的文件对应的
	r.GET("/news", func(context *gin.Context) {
		context.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是一个后台数据",
		})

	})
	r.Run()
}
