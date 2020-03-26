package yna

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/yna")
	{
		posts.POST("/yna_by_adunit", PostYNAReport)
	}
}
