package route

import (
	"gin_work/controller"
	"github.com/gin-gonic/gin"
)

type Route struct{}

func (*Route) Login(r *gin.Engine) {
	var login *controller.Login
	r.POST("/user_login", login.UserLogin)
}
