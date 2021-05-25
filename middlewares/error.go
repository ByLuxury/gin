package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

// ErrorHandler 通用错误处理
func ErrorHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", err)
			debug.PrintStack()
			if ctx.Writer.Status() != http.StatusOK {
				ctx.Status(http.StatusOK)
			}
			ctx.JSON(http.StatusOK, err)
		}
	}()
	ctx.Next()
}
