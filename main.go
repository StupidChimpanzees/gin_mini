package main

import (
	"gin_work/wrap/config"
	"gin_work/wrap/cookie"
	"gin_work/wrap/database"
	log2 "gin_work/wrap/log"
	"gin_work/wrap/middleware"
	"gin_work/wrap/preload"
	"gin_work/wrap/route"
	"gin_work/wrap/session"
	"gin_work/wrap/validator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := handler()

	preload.Load()

	// autotls.Run(r, "example1.com")
	err := r.Run(":" + strconv.Itoa(config.Mapping.App.Port))
	if err != nil {
		log2.Error(err.Error())
		log.Fatalf("error info: " + err.Error())
	}
}

func handler() *gin.Engine {
	r := gin.Default()

	// 加载配置
	err := config.Load("config.yaml")
	if err != nil {
		log2.Error(err.Error())
		panic(err.Error())
	}

	// 加载全局中间件
	middleware.Load(r)

	// 加载cookie和session配置
	cookie.Load()
	store := session.Load()
	r.Use(sessions.Sessions("GlobalSession", store))

	// 设置database
	database.SetDbEngine()

	// 加载view配置
	// 目录下必须有.html文件才能使用
	// view.Load(r)

	// 注册自定义验证
	validator.Load()

	// 构建路由
	route.Load(r)

	_ = r.SetTrustedProxies(nil)

	return r
}
