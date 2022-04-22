package baseContrllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseControllers struct {
}

func (con BaseControllers) Success(r *gin.Context) {
	//这个地方是在终端显示
	username, _ := r.Get("username")
	fmt.Println(username)
	//下面这个是类型断言，是在网页上显示
	v, ok := username.(string)
	if ok == false {
		fmt.Println("err")
	} else {
		r.String(http.StatusOK, "成功"+v)
	}

}
func (con BaseControllers) Fail(r2 *gin.Context) {
	r2.String(http.StatusOK, "失败")
}
