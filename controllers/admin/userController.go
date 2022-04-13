package admin

import (
	"123/baseContrllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	baseContrllers.BaseControllers
}

func (con UserController) Password(c *gin.Context) {
	//c.String(http.StatusOK, "我是admin--password--我是func分装的")
	con.Success(c)
}
func (con UserController) Age(c *gin.Context) {
	c.String(http.StatusOK, "我是admin--age-------我是func分装的")
	con.Fail(c)
}
