package session

import (
	"gin_work/wrap/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Load() cookie.Store {
	return cookie.NewStore([]byte(getConfig().Secret))
}

func getConfig() config.SessionConfiguration {
	return config.Mapping.Session
}

func Set(context *gin.Context, name string, value interface{}) error {
	session := sessions.Default(context)
	session.Options(sessions.Options{MaxAge: getConfig().Expire})
	session.Set(name, value)
	return session.Save()
}

func Get(context *gin.Context, name string) interface{} {
	session := sessions.Default(context)
	return session.Get(name)
}

func Delete(context *gin.Context, name string) {
	session := sessions.Default(context)
	session.Delete(name)
}

func Clear(context *gin.Context) {
	session := sessions.Default(context)
	session.Clear()
}
