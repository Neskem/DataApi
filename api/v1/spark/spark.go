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
		posts.GET("/pv_by_author/:start_date/:end_date/:author", readAuthorPV)
		posts.GET("/pv_by_hostname/:start_date/:end_date/:hostname", readHostPV)
	}
}
