package spark

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/spark")
	{
		posts.POST("/pv_by_urls", ReadDailyPV)
		posts.POST("/pv_monthly_by_urls", ReadMonthlyPV)
		posts.POST("total_pv_by_urls", ReadTotalPV)
		posts.GET("/pv_by_author", GetAuthorPV)
		posts.GET("/pv_by_hostname", GetHostPV2)
	}
}
