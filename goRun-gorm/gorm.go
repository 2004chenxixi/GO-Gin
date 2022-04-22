package main

import (
	"123/gorm"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/p", func(c *gin.Context) {
		//查询资源库
		bankSlice := []gorm.Bank{} //这里用到了gorm，gorm.Bank这个是已经定义好的，结构体在gorm中
		gorm.DB.Find(&bankSlice)
		fmt.Println(bankSlice)
		c.JSON(http.StatusOK, gin.H{
			"result": bankSlice,
		})
	})

	r.GET("/", func(c *gin.Context) {
		//钱大于100--查询
		bankSlice := []gorm.Bank{} //这里用到了gorm，gorm.Bank这个是已经定义好的，结构体在gorm中
		gorm.DB.Where("money>100").Find(&bankSlice)
		fmt.Println(bankSlice)
		c.JSON(http.StatusOK, gin.H{
			"result": bankSlice,
		})
	})

	r.GET("/x", func(c *gin.Context) {
		//在已经创建的brank下面，新建资料，增加数据
		bankSlice := gorm.Bank{ //这里不是空切片不需要[]，定义数据
			Id:       3,
			Username: "林平之",
			Money:    12213,
		} //这里用到了gorm，gorm.Bank这个是已经定义好的，结构体在gorm中
		gorm.DB.Create(&bankSlice) //因为是新建，所以用Create
		fmt.Println(bankSlice)
		c.JSON(http.StatusOK, gin.H{
			"result": bankSlice,
		})
	})

	r.GET("/w", func(c *gin.Context) {
		//修改数据
		//查询ID=3的数据
		bankSlice := gorm.Bank{Username: "我不是林平之"} //这里用到了gorm，gorm.Bank这个是已经定义好的，结构体在gorm中
		gorm.DB.Find(&bankSlice)                   //查询到id=3的数据，把它保存，下面可以进行修改
		//对ID=3，的数据进行修改+保存
		//	bankSlice.Username = "我不之"
		bankSlice.Money = 10
		gorm.DB.Save(bankSlice)
		c.JSON(http.StatusOK, gin.H{
			"result": bankSlice,
		})
	})
	r.GET("/e", func(c *gin.Context) {
		//删除
		bankSlice := gorm.Bank{Id: 3} //这里用到了gorm，gorm.Bank这个是已经定义好的，结构体在gorm中
		gorm.DB.Delete(&bankSlice)    //查询到id=3的数据，删除
		c.JSON(http.StatusOK, gin.H{
			"result": bankSlice,
		})
	})

	r.Run(":909") // listen and serve on 0.0.0.0:8080
}

//gorm/gorm.go
//gorm/bank
//分装在这两个包里
