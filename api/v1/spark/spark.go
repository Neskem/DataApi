package spark

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/spark")
	{
		posts.POST("/pv_by_urls", readDailyPV)
		posts.POST("/pv_monthly_by_urls", readMonthlyPV)
		posts.POST("total_pv_by_urls", readTotalPV)
		posts.GET("/pv_by_author", getAuthorPV)
		posts.GET("/pv_by_hostname", getHostPV)
	}
}
