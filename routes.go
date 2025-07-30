package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up all routes and returns the Gin engine
func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/students", CreateStudent(db))
		api.GET("/students/:id", GetStudent(db))
		api.PUT("/students/:id", UpdateStudent(db))
		api.DELETE("/students/:id", DeleteStudent(db))
		// Add other API routes here as needed
	}

	r.GET("/healthcheck", HealthCheck)

	// You can add /metrics route here if using Prometheus

	return r
}
