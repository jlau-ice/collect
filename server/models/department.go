package models

import (
	"time"

	"gorm.io/gorm"
)

// Department 部门模型
type Department struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name        string `gorm:"type:varchar(100);not null;uniqueIndex" json:"name"` // 部门名称
	Description string `gorm:"type:varchar(500)" json:"description"`                // 部门描述

	// 关联关系
	Users []User `gorm:"foreignKey:DepartmentID" json:"users,omitempty"`
}

// TableName 指定表名
func (Department) TableName() string {
	return "departments"
}

