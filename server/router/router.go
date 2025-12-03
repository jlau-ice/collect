package router

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine) {
	// API路由组
	api := r.Group("/api")
	{
		// 健康检查
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
				"message": "服务运行正常",
			})
		})

		// 部门管理
		departments := api.Group("/departments")
		{
			departments.GET("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "获取部门列表"})
			})
			departments.POST("", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "创建部门"})
			})
			departments.PUT("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "更新部门"})
			})
			departments.DELETE("/:id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "删除部门"})
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

