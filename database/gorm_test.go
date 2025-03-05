package database

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T) {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接出错：", err)
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(100)          // 空闲连接数
	sqlDB.SetMaxOpenConns(1000)         // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接生命周期
}
