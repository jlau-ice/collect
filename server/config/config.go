package config

import (
	"fmt"
	"os"
	"strconv"

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
	SSLMode  string
	TimeZone string
	Schema   string
}

// UploadConfig 上传配置
type UploadConfig struct {
	BasePath string // 文件存储基础路径
	MaxSize  int64  // 最大文件大小（字节）
}

var AppConfig *Config

// LoadConfig 加载配置
// LoadConfig 加载配置
func LoadConfig() error {
	// 尝试加载.env文件（如果存在）
	_ = godotenv.Load()
	// 默认文件最大大小（10MB）
	defaultMaxSize := int64(10 * 1024 * 1024)
	// 从环境变量获取 MaxSize，如果获取失败则使用默认值
	maxSizeEnv := getEnv("UPLOAD_MAX_SIZE", strconv.FormatInt(defaultMaxSize, 10))
	maxSize, err := strconv.ParseInt(maxSizeEnv, 10, 64)
	if err != nil {
		fmt.Printf("Warning: UPLOAD_MAX_SIZE is invalid, using default %d bytes\n", defaultMaxSize)
		maxSize = defaultMaxSize
	}
	AppConfig = &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			// PgSQL 默认端口为 5432
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "collect"),
			// PgSQL 连接参数
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
			TimeZone: getEnv("DB_TIME_ZONE", "Asia/Shanghai"),
			Schema:   getEnv("DB_SCHEMA", "public"),
		},
		Upload: UploadConfig{
			BasePath: getEnv("UPLOAD_BASE_PATH", "./uploads"),
			MaxSize:  maxSize, // 使用解析后的值
		},
	}

	return nil
}

// GetDSN 获取数据库连接字符串
// 修复后的 GetDSN 获取 PostgreSQL 数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	// 使用标准的 key=value 格式，不需要额外的 'options=' 关键字包裹。
	// search_path 应该直接作为 DSN 的一个参数。
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s search_path=%s",
		c.Host,
		c.User,
		c.Password,
		c.DBName,
		c.Port,
		c.SSLMode,
		c.TimeZone,
		c.Schema,
	)
}
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
