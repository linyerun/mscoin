package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func gormMysqlInit(dsn string) *gorm.DB {
	var err error

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("连接数据库失败, error=" + err.Error())
	}

	db, err := _db.DB()
	if err != nil {
		log.Fatal("获取DB()失败, error=" + err.Error())
	}

	// 连接池配置（配置数据库连接池的设置是持久有效的，因此在调用_db.DB()时，配置仍然存在）
	db.SetMaxOpenConns(100) // 最大连接数
	db.SetMaxIdleConns(10)  // 最大空闲连接数

	return _db
}
