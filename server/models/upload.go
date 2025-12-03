package models

import (
	"time"

	"gorm.io/gorm"
)

// Upload 上传记录模型
type Upload struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	UserID     uint   `gorm:"not null;index" json:"user_id"`     // 上传人员ID
	TaskID     uint   `gorm:"not null;index" json:"task_id"`     // 任务ID
	FileName   string `gorm:"type:varchar(255);not null" json:"file_name"` // 文件名
	FilePath   string `gorm:"type:varchar(500);not null" json:"file_path"` // 文件存储路径
	FileSize   int64  `gorm:"not null" json:"file_size"`         // 文件大小（字节）
	MimeType   string `gorm:"type:varchar(100)" json:"mime_type"` // MIME类型

	// 关联关系
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Task Task `gorm:"foreignKey:TaskID" json:"task,omitempty"`
}

// TableName 指定表名
func (Upload) TableName() string {
	return "uploads"
}

