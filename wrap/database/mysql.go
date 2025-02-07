package database

import (
	"fmt"
	log2 "gin_work/wrap/log"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConf struct{}

func (*MysqlConf) dsn(username string, password string, host string, port int, dbname string, charset string,
	parseTime bool, local string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, dbname, charset, parseTime, local)
}

func (m *MysqlConf) SetDb() (string, *gorm.Config) {
	var conf gorm.Config

	dsn := m.dsn(DBConfig.Username, DBConfig.Password, DBConfig.Host,
		DBConfig.Port, DBConfig.DBName, DBConfig.Charset, DBConfig.ParseTime, DBConfig.Loc)
	SetDbLog(&conf)

	return dsn, &conf
}

func (m *MysqlConf) Open(dsn string, conf *gorm.Config) {
	dialectal := mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	})
	dbInstance, err := gorm.Open(dialectal, conf)
	if err != nil {
		log2.Error(err.Error())
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	DB = dbInstance
}
