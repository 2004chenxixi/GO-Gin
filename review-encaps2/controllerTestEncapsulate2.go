package review_encaps2

import (
	controllerTestEncapsulate "123/review-encapsulate"
	"github.com/gin-gonic/gin"
)

func GainSumGet(c *gin.Engine) {
	c.GET("/user2", controllerTestEncapsulate.Controller{}.EncapsulateGet)
}
func GainSumPost(r *gin.Engine) {
	r.POST("/DoAddUser", controllerTestEncapsulate.Controller{}.EncapsulatePost)
}
