package baseContrllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseControllers struct {
}

func (con BaseControllers) Success(r *gin.Context) {
	r.String(http.StatusOK, "成功")
}
func (con BaseControllers) Fail(r2 *gin.Context) {
	r2.String(http.StatusOK, "失败")
}
