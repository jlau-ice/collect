package service

import (
	"context"
	"errors"

	"github.com/jlau-ice/collect/internal/models"
	"github.com/jlau-ice/collect/internal/types/interfaces"
)

// DepartmentServiceImpl 部门业务实现
type DepartmentServiceImpl struct {
	repo interfaces.DepartmentRepository
}

// NewDepartmentService 创建部门业务实例
func NewDepartmentService(repo interfaces.DepartmentRepository) *DepartmentServiceImpl {
	return &DepartmentServiceImpl{repo: repo}
}

// ListDepartments 获取部门列表
func (s *DepartmentServiceImpl) ListDepartments(ctx context.Context) ([]models.Department, error) {
	return s.repo.List(ctx)
}

// CreateDepartment 创建部门
func (s *DepartmentServiceImpl) CreateDepartment(ctx context.Context, name, desc string) (*models.Department, error) {
	if name == "" {
		return nil, errors.New("部门名称不能为空")
	}
	dept := &models.Department{
		Name:        name,
		Description: desc,
	}
	if err := s.repo.Create(ctx, dept); err != nil {
		return nil, err
	}
	return dept, nil
}
