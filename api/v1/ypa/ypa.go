package ypa

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/ypa")
	{
		posts.GET("/daily", getDailyYPA)
	}
}
