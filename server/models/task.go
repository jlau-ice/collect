package models

import (
	"time"

	"gorm.io/gorm"
)

// Task 转发任务模型
type Task struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string    `gorm:"type:varchar(200);not null" json:"name"` // 转发视频/文章名称
	WeekStart   time.Time `gorm:"not null;index" json:"week_start"`      // 周开始日期
	WeekEnd     time.Time `gorm:"not null;index" json:"week_end"`        // 周结束日期
	WeekNumber  int       `gorm:"not null;index" json:"week_number"`      // 第几周
	Description string    `gorm:"type:varchar(500)" json:"description"`  // 任务描述

	// 关联关系
	Uploads []Upload `gorm:"foreignKey:TaskID" json:"uploads,omitempty"`
}

// TableName 指定表名
func (Task) TableName() string {
	return "tasks"
}

// GetWeekRange 获取周日期区间字符串，格式：2025-11-24～2025-11-30
func (t *Task) GetWeekRange() string {
	return t.WeekStart.Format("2006-01-02") + "～" + t.WeekEnd.Format("2006-01-02")
}

