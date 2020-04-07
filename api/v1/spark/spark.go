package spark

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/spark")
	{
		posts.POST("/pv_by_urls", PostDailyPV)
		posts.POST("/pv_monthly_by_urls", PostMonthlyPV)
		posts.POST("/total_pv_by_urls", PostTotalPV)
		posts.POST("/move_zi_article_pv", PostMovePV)
		posts.GET("/pv_by_author", GetAuthorPV)
		posts.GET("/pv_by_hostname", GetHostPV)
	}
}
