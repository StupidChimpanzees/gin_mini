package middleware

import (
	"errors"
	"fmt"
	"gin_work/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GlobalMiddleware struct{}

func (*GlobalMiddleware) Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	}
}

// Cors 开启跨域请求
func (*GlobalMiddleware) Cors() gin.HandlerFunc {
	return cors.Default()
}

// 全局错误处理中间件
func (*GlobalMiddleware) Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch v := err.(type) {
				case error:
					common.Error(err.(error), c, nil)
				case string:
					common.Error(errors.New(err.(string)), c, nil)
				default:
					fmt.Println(v)
				}
			}
		}()
		c.Next()
	}
}
