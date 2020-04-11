package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandler 通用错误处理
func ErrorHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.Writer.Status() != http.StatusOK {
				ctx.Status(http.StatusOK)
			}
				ctx.JSON(http.StatusOK, err)
 		}
	}()
	ctx.Next()
}
