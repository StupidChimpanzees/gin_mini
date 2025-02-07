package go_test

import (
	"gin_work/wrap/config"
	"gin_work/wrap/cookie"
	"gin_work/wrap/log"
	"gin_work/wrap/middleware"
	"gin_work/wrap/route"
	"gin_work/wrap/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func instance() *gin.Engine {
	r := gin.Default()

	// 加载配置
	err := config.Load("../config.yaml")
	if err != nil {
		log.Error(err.Error())
		panic("Config file load error")
	}

	// 加载全局中间件
	middleware.Load(r)

	// 加载cookie和session配置
	cookie.Load()
	store := session.Load()
	r.Use(sessions.Sessions("GlobalSession", store))

	// 加载view配置
	// 目录下必须有.html文件才能使用
	// view.Load(r)

	// 构建路由
	route.Load(r)

	return r
}

func TestIndexRoute(t *testing.T) {
	router := instance()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/index", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUserLogin(t *testing.T) {
	router := instance()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func Benchmark_IndexRoute(b *testing.B) {
	b.StopTimer()
	router := instance()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/index", nil)
		router.ServeHTTP(w, req)
	}
}
