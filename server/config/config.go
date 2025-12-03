package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config 应用配置
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Upload   UploadConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
	Mode string // debug, release, test
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Charset  string
}

// UploadConfig 上传配置
type UploadConfig struct {
	BasePath string // 文件存储基础路径
	MaxSize  int64  // 最大文件大小（字节）
}

var AppConfig *Config

// LoadConfig 加载配置
func LoadConfig() error {
	// 尝试加载.env文件（如果存在）
	_ = godotenv.Load()

	AppConfig = &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "collect_tools"),
			Charset:  getEnv("DB_CHARSET", "utf8mb4"),
		},
		Upload: UploadConfig{
			BasePath: getEnv("UPLOAD_BASE_PATH", "./uploads"),
			MaxSize:  10 * 1024 * 1024, // 默认10MB
		},
	}

	return nil
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.Charset)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

