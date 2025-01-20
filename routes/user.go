package routes

import (
	"connect-api/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes กำหนด Routes ที่เกี่ยวข้องกับผู้ใช้งาน
func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("", controllers.GetUsers)    // ดึงข้อมูลผู้ใช้งาน
		userGroup.POST("", controllers.CreateUser) // เพิ่มผู้ใช้งานใหม่
	}
}
