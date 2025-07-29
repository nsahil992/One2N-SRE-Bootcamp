package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Attach Prometheus middleware globally
	r.Use(PrometheusMiddleware())

	api := r.Group("/api/v1")
	{
		api.POST("/students", CreateStudent(db))
		api.GET("/students", GetAllStudents(db))
		api.GET("/students/:id", GetStudent(db))
		api.PUT("/students/:id", UpdateStudent(db))
		api.DELETE("/students/:id", DeleteStudent(db))
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Student API"})
	})

	r.GET("/healthcheck", HealthCheck)

	// Serve Prometheus metrics
	r.GET("/metrics", MetricsHandler())

	return r
}
