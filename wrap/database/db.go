package database

import (
	"gin_work/wrap/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbInterface interface {
	dsn(username string, password string, host string, port int,
		dbname string, charset string, parseTime bool, local string) string
	Open(dsn string, conf *gorm.Config)
}

type dbConfig struct {
	DBType    string
	DBName    string
	Username  string
	Password  string
	Host      string
	Port      int
	Charset   string
	ParseTime bool
	Loc       string
}

var (
	DBConfig *dbConfig
	DB       *gorm.DB
)

func init() {
	DBConfig = NewDBConfig()
}

func NewDBConfig() *dbConfig {
	dbc := config.Mapping.Database
	return &dbConfig{
		DBType:    dbc.DBType,
		DBName:    dbc.Name,
		Username:  dbc.Username,
		Password:  dbc.Password,
		Host:      dbc.Host,
		Port:      dbc.Port,
		Charset:   dbc.Charset,
		ParseTime: true,
		Loc:       "Local",
	}
}

func SetDbEngine() {
	if DBConfig.DBType == "mysql" {
		if poolConf.Enable {
			poolConf.Open(poolConf.SetPool())
		} else {
			var mysqlConf *MysqlConf
			mysqlConf.Open(mysqlConf.SetDb())
		}
	}
}

func SetDbLog(conf *gorm.Config) {
	conf.Logger = logger.Default.LogMode(logger.Info)
}
