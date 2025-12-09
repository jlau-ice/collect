package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlau-ice/collect/internal/service"
)

// DepartmentHandler 部门控制器
type DepartmentHandler struct {
	svc *service.DepartmentServiceImpl
}

// NewDepartmentHandler 创建控制器
func NewDepartmentHandler(svc *service.DepartmentServiceImpl) *DepartmentHandler {
	return &DepartmentHandler{svc: svc}
}

// List 部门列表
func (h *DepartmentHandler) List(c *gin.Context) {
	depts, err := h.svc.ListDepartments(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": depts})
}

// Create 创建部门
func (h *DepartmentHandler) Create(c *gin.Context) {
	var req struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	dept, err := h.svc.CreateDepartment(c.Request.Context(), req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dept})
}
