package dbinit

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateGormMysqlClient(dsn string, maxOpenConns, maxIdleConns int) (*gorm.DB, error) {
	// 连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	// 连接池配置（配置数据库连接池的设置是持久有效的，因此在调用_db.DB()时，配置仍然存在）
	_db, err := db.DB()
	if err != nil {
		return nil, err
	}
	_db.SetMaxOpenConns(maxOpenConns) // 最大连接数
	_db.SetMaxIdleConns(maxIdleConns) // 最大空闲连接数

	return db, nil
}
