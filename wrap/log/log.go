package log

import (
	"fmt"
	"gin_work/wrap/config"
	"log"
	"os"
	"strings"
	"time"
)

type appLog struct {
	Level        int
	LevelStr     string
	Path         string
	Prefix       string
	Flags        int
	ConsolePrint bool
}

const (
	LevelError = iota
	LevelWarning
	LevelInfo
)

var (
	AppLog *appLog
)

func init() {
	AppLog = getConfig()
}

func getConfig() *appLog {
	conf := config.Mapping.Log
	return &appLog{
		Level:        AppLog.getLevel(conf.Level),
		LevelStr:     conf.Level,
		Path:         conf.Path,
		Prefix:       "",
		Flags:        AppLog.getFlag(conf.Format),
		ConsolePrint: conf.ConsolePrint,
	}
}

func New(level string) (*log.Logger, *os.File) {
	err := os.MkdirAll(AppLog.getPath(AppLog.Path), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(AppLog.getPath(AppLog.Path)+level+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return log.New(file, strings.ToUpper(level)+": ", AppLog.Flags), file
}

func Error(s ...any) {
	Recode("error", s...)
}

func Warning(s ...any) {
	Recode("warning", s...)
}

func Info(s ...any) {
	Recode("info", s...)
}

func Write(s ...any) {
	Recode(AppLog.LevelStr, s...)
}

func Recode(level string, s ...any) {
	logObj, file := New(level)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(file)
	logObj.Println(s)
	if AppLog.ConsolePrint {
		fmt.Println(s)
	}
}

func (a *appLog) getPath(path string) string {
	path = strings.Replace(path, "{date}", time.Now().Format("2006-01-02"), -1)
	path = strings.Replace(path, "{level}", strings.ToLower(a.LevelStr), -1)
	return path
}

func (a *appLog) getFlag(format string) int {
	f := strings.Split(format, "|")
	var flags int
	for _, s := range f {
		switch strings.ToLower(s) {
		case "date":
			flags = flags | log.Ldate
		case "time":
			flags = flags | log.Ltime
		case "longfile":
			flags = flags | log.Llongfile
		case "shortfile":
			flags = flags | log.Lshortfile
		case "utc":
			flags = flags | log.LUTC
		}
	}
	return flags
}

func (a *appLog) getLevel(l string) int {
	switch strings.ToLower(l) {
	case "error":
		return LevelError
	case "warning":
		return LevelWarning
	default:
		return LevelInfo
	}
}
