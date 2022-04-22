package main

import (
	//"123/Routers"
	"123/Routers"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	Routers.ApiRouters(r)
	Routers.AdminRouters(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}

//baseContrllers/base.go
//controllers/admin/userController.go
//Routers/adminRouters.go
//Gin-06-pull-away/main-inherit.go
//依此向下，结构体的继承  ---包名+方法名
