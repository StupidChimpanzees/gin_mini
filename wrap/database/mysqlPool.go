package database

import (
	"database/sql"
	"gin_work/wrap/config"
	log2 "gin_work/wrap/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type PoolConf struct {
	Enable          bool
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

var poolConf *PoolConf

func init() {
	poolConf = NewPoolConf()
}

func NewPoolConf() *PoolConf {
	conf := &PoolConf{
		Enable:          true,
		MaxIdleConn:     10,
		MaxOpenConn:     100,
		ConnMaxIdleTime: 3 * time.Second,
		ConnMaxLifetime: 180 * time.Second,
	}
	pConf := conf.getConf()
	conf.Enable = pConf.Enable
	if pConf.MaxIdleConn != 0 {
		conf.MaxIdleConn = pConf.MaxIdleConn
	}
	if pConf.MaxOpenConn != 0 {
		conf.MaxOpenConn = pConf.MaxOpenConn
	}
	if pConf.ConnMaxIdleTime != 0 {
		conf.ConnMaxIdleTime = time.Duration(pConf.ConnMaxIdleTime) * time.Second
	}
	if pConf.ConnMaxLifeTime != 0 {
		conf.ConnMaxLifetime = time.Duration(pConf.ConnMaxLifeTime) * time.Second
	}
	return conf
}

func (*PoolConf) getConf() config.DatabasePoolConfiguration {
	return config.Mapping.Database.Pool
}

func (*PoolConf) SetPool() (*sql.DB, *gorm.Config) {
	var mysqlConf *MysqlConf
	dsn, conf := mysqlConf.SetDb()

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log2.Error(err.Error())
		panic(err)
	}

	sqlDB.SetMaxIdleConns(poolConf.MaxIdleConn)
	sqlDB.SetMaxOpenConns(poolConf.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(poolConf.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(poolConf.ConnMaxLifetime)

	return sqlDB, conf
}

func (*PoolConf) Open(sqlDb *sql.DB, conf *gorm.Config) {
	dialectal := mysql.New(mysql.Config{
		Conn:                      sqlDb, // DSN data source name
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
