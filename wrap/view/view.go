package view

import (
	"gin_work/wrap/config"
	"github.com/gin-gonic/gin"
)

type VConfig struct {
	TempPath   string
	StaticPath string
	DelimBegin string
	DelimEnd   string
}

func Load(r *gin.Engine) {
	vc := NewViewConfig()
	vc.setTempPath(r, vc.TempPath)
	vc.setStaticPath(r, vc.StaticPath)
	vc.setDelim(r, vc.DelimBegin, vc.DelimEnd)
}

func NewViewConfig() *VConfig {
	vc := &VConfig{
		TempPath:   "public/**/*",
		StaticPath: "/assets",
		DelimBegin: "{{",
		DelimEnd:   "}}",
	}
	globalConf := config.Mapping.View
	if globalConf.TempPath != "" {
		vc.TempPath = globalConf.TempPath
	}
	if globalConf.StaticPath != "" {
		vc.StaticPath = globalConf.StaticPath
	}
	if globalConf.DelimBegin != "" {
		vc.DelimBegin = globalConf.DelimBegin
	}
	if globalConf.DelimEnd != "" {
		vc.DelimEnd = globalConf.DelimEnd
	}
	return vc
}

func (*VConfig) setTempPath(r *gin.Engine, path string) {
	r.LoadHTMLGlob(path)
}

func (*VConfig) setStaticPath(r *gin.Engine, path string) {
	r.Static("/assets", path)
}

func (*VConfig) setDelim(r *gin.Engine, begin, end string) {
	r.Delims(begin, end)
}
