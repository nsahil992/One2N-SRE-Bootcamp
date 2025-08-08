package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up all routes and returns the Gin engine
func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Ngineers, from One2N Bootcamp ⬆️"})
	})

	api := r.Group("/api/v1")
	{
		api.POST("/students", CreateStudent(db))
		api.GET("/students", GetAllStudents(db))
		api.GET("/students/:id", GetStudent(db))
		api.PUT("/students/:id", UpdateStudent(db))
		api.DELETE("/students/:id", DeleteStudent(db))
	}

	r.GET("/healthcheck", HealthCheck)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.GET("/metrics", MetricsHandler()) // Prometheus metrics route

	return r
}
