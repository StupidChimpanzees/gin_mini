package cookie

import (
	"gin_work/wrap/config"

	"github.com/gin-gonic/gin"
)

type cookiePack struct {
	Expire   int
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

var cPack *cookiePack

func Load() {
	cPack = getConfig()
}

func getConfig() *cookiePack {
	conf := config.Mapping.Cookie
	return &cookiePack{
		Expire:   conf.Expire,
		Path:     conf.Path,
		Domain:   conf.Domain,
		Secure:   conf.Secure,
		HttpOnly: conf.HttpOnly,
	}
}

func Set(c *gin.Context, name, value string) {
	c.SetCookie(name, value, cPack.Expire, cPack.Path, cPack.Domain, cPack.Secure, cPack.HttpOnly)
}

func Get(c *gin.Context, name string) (string, error) {
	return c.Cookie(name)
}

func Delete(c *gin.Context, name string) {
	c.SetCookie(name, "", -1, cPack.Path, cPack.Domain, cPack.Secure, cPack.HttpOnly)
}
