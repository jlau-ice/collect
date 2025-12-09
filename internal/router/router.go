package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jlau-ice/collect/internal/handler"
	"go.uber.org/dig"
)

type RouterParams struct {
	dig.In
	AuthHandler *handler.AuthHandler
}

func NewRouter(params RouterParams) *gin.Engine {
	r := gin.New()

	// CORS 中间件应放在最前面
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-API-Key", "X-Request-ID"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	v1 := r.Group("/api/v1")
	{
		RegisterUserRoutes(v1, params.AuthHandler)
	}
	return r
}

func RegisterUserRoutes(r *gin.RouterGroup, handler *handler.AuthHandler) {
	r.POST("/auth/register", handler.AddUser)
}

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine, deptHandler *handler.DepartmentHandler) {
	// API路由组
	api := r.Group("/api")
	{
		// 健康检查
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "服务运行正常",
			})
		})

		// 部门管理
		departments := api.Group("/departments")
		{
			departments.GET("", deptHandler.List)
			departments.POST("", deptHandler.Create)
			departments.PUT("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "更新部门（示例待实现）"})
			})
			departments.DELETE("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "删除部门（示例待实现）"})
			})
		}

		// 人员管理
		users := api.Group("/users")
		{
			users.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "获取人员列表"})
			})
			users.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "创建人员"})
			})
			users.PUT("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "更新人员"})
			})
			users.DELETE("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "删除人员"})
			})
		}

		// 转发任务
		tasks := api.Group("/tasks")
		{
			tasks.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "获取任务列表"})
			})
			tasks.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "创建任务"})
			})
			tasks.PUT("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "更新任务"})
			})
			tasks.DELETE("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "删除任务"})
			})
		}

		// 文件上传
		upload := api.Group("/upload")
		{
			upload.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "上传文件"})
			})
		}
	}
}
