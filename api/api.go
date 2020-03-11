package api

import (
	apiV1 "DataApi.Go/api/v1"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		apiV1.ApplyRoutes(api)
	}
}
