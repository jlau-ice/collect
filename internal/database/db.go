package database

import (
	"log"

	"github.com/jlau-ice/collect/internal/config"
	"github.com/jlau-ice/collect/internal/models"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	dsn := config.AppConfig.Database.GetDSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// GORM 日志配置
		Logger: logger.Default.LogMode(logger.Info),

		// ❗ 注意：如果您在 config 包中设置了 Schema，
		// 这里 GORM 默认就会使用该 Schema 进行操作。
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
	// PgSQL 对大小写敏感，且习惯使用蛇形命名 (snake_case)。
	// GORM 默认会处理这些，但请确保你的 models/ 结构体命名规范。
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
