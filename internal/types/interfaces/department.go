package interfaces

import (
	"context"

	"github.com/jlau-ice/collect/internal/models"
)

// DepartmentRepository 定义部门数据访问接口
type DepartmentRepository interface {
	List(ctx context.Context) ([]models.Department, error)
	Create(ctx context.Context, dept *models.Department) error
}

// DepartmentService 定义部门业务接口
type DepartmentService interface {
	ListDepartments(ctx context.Context) ([]models.Department, error)
	CreateDepartment(ctx context.Context, name, desc string) (*models.Department, error)
}
