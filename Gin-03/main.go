package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type criminal struct {
	Person string
	Gender string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(context *gin.Context) {

		context.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",

			"score": 89,

			"beautiful": []string{"方圆", "芬芳", "小野"},

			"criminalSlice": []interface{}{
				&criminal{
					Person: "李四",
					Gender: "男",
				},
			},

			"randomSlice": []string{},

			"new": &criminal{
				Person: "赵五",
				Gender: "不男不女",
			},
		})
	})

	r.GET("/news", func(context *gin.Context) {
		news := &criminal{
			Person: "张三",
			Gender: "男",
		}
		context.HTML(http.StatusOK, "default/news.html", gin.H{
			//这里的name写的是需要导入的部分，写全了
			//name还要配套r.LoadHTMLGlob("templates/**/*")来用
			//**/表示一个目录
			"title": "罪犯信息",
			"news":  news,
		})
	})
	r.Run()
}
