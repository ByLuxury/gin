package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"fmt"
	"net/http"
)

func Cros() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			method := ctx.Request.Method
			origin := ctx.Request.Header.Get("Origin")
			var headerKeys []string
			for k, _ := range ctx.Request.Header {
				headerKeys = append(headerKeys, k)
			}
			headerStr := strings.Join(headerKeys, ", ")
			if headerStr != "" {
				headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
			} else {
				headerStr = "access-control-allow-origin, access-control-allow-headers"
			}
			if origin != "" {
				ctx.Header("Access-Control-Allow-Origin", "*")
				ctx.Header("Access-Control-Allow-Headers", headerStr)
				ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
				ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
				ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
				ctx.Set("content-type", "application/json")
			}
			if method == "OPTIONS" {
				ctx.JSON(http.StatusOK, "Options Request!")
			}
			ctx.Next()
		}
}