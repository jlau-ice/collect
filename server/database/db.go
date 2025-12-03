package database

import (
	"log"
	"server/config"
	"server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	dsn := config.AppConfig.Database.GetDSN()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	log.Println("数据库连接成功")

	// 自动迁移
	err = AutoMigrate()
	if err != nil {
		return err
	}

	return nil
}

// AutoMigrate 自动迁移数据库表
func AutoMigrate() error {
	err := DB.AutoMigrate(
		&models.Department{},
		&models.User{},
		&models.Task{},
		&models.Upload{},
	)

	if err != nil {
		return err
	}

	log.Println("数据库表迁移完成")
	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

