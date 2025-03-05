package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库配置结构体
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// InitializeDB 完整的数据库初始化方法
func InitializeDB(config *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库实例失败: %w", err)
	}

	sqlDB.SetMaxOpenConns(100) // 最大打开连接数
	sqlDB.SetMaxIdleConns(10)  // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("数据库心跳检测失败: %w", err)
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, fmt.Errorf("表结构迁移失败: %w", err)
	}

	fmt.Println("数据库连接成功！")
	return db, nil
}
