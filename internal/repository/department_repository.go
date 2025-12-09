package repository

import (
	"context"

	"github.com/jlau-ice/collect/internal/models"
	"github.com/jlau-ice/collect/internal/types/interfaces"
	"gorm.io/gorm"
)

// DepartmentRepositoryImpl 部门仓储实现
type DepartmentRepositoryImpl struct {
	db *gorm.DB
}

// NewDepartmentRepository 创建仓储实例
func NewDepartmentRepository(db *gorm.DB) interfaces.DepartmentRepository {
	return &DepartmentRepositoryImpl{db: db}
}

// List 返回部门列表
func (r *DepartmentRepositoryImpl) List(ctx context.Context) ([]models.Department, error) {
	var depts []models.Department
	if err := r.db.WithContext(ctx).Find(&depts).Error; err != nil {
		return nil, err
	}
	return depts, nil
}

// Create 新增部门
func (r *DepartmentRepositoryImpl) Create(ctx context.Context, dept *models.Department) error {
	return r.db.WithContext(ctx).Create(dept).Error
}
