package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB 初始化数据库连接
func InitDB(dsn string, debug bool) error {
	// 使用 MySQL 驱动连接数据库
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// 如果需要开启调试模式
	if debug {
		db = db.Debug()
	}

	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return db
}
