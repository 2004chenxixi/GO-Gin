package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Println(str01 string, str02 string) string {
	return str01 + str02
}
func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"Println": Println,
	})

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/gin04test.html", gin.H{
			"m": "chen",
			"b": "xixi",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
