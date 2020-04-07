package adSense

import (
"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/adsense")
	{
		posts.POST("/daily", PostDailyAdSense)
		posts.POST("/adexchange/revenue", PostDailyAdSenseRevenue)
		posts.GET("/domains", GetAdSenseDomains)
	}
}

