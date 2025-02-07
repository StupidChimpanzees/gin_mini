package common

import (
	"encoding/json"
	"gin_work/wrap/log"
	"github.com/gin-gonic/gin"
)

func Log(level string, err error, c *gin.Context, params any) {
	b, _ := json.Marshal(params)
	log.Recode(level, c.Request.Method+" "+c.Request.URL.RequestURI()+" : "+string(b))
	if err != nil {
		log.Recode(level, err.Error())
	}
}

func Error(err error, c *gin.Context, params any) {
	Log("error", err, c, params)
}

func Warning(err error, c *gin.Context, params any) {
	Log("warning", err, c, params)
}

func Info(c *gin.Context, params any) {
	Log("info", nil, c, params)
}
