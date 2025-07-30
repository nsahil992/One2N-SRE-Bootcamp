package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns simple OK status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// CreateStudent handles POST /students
func CreateStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var student Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var id int
		err := db.QueryRow(`INSERT INTO students (name, age, email) VALUES ($1, $2, $3) RETURNING id`,
			student.Name, student.Age, student.Email).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		student.ID = id
		c.JSON(http.StatusCreated, student)
	}
}

// GetStudent handles GET /students/:id
func GetStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		var s Student

		err := db.QueryRow(`SELECT id, name, age, email FROM students WHERE id = $1`, idParam).
			Scan(&s.ID, &s.Name, &s.Age, &s.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, s)
	}
}

// UpdateStudent handles PUT /students/:id
func UpdateStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		var student Student

		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := db.Exec(`UPDATE students SET name=$1, age=$2, email=$3 WHERE id=$4`,
			student.Name, student.Age, student.Email, idParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student updated"})
	}
}

// DeleteStudent handles DELETE /students/:id
func DeleteStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		res, err := db.Exec(`DELETE FROM students WHERE id = $1`, idParam)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
	}
}

// Example function showing query with rows.Close() error handling
func GetAllStudents(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, age, email FROM students")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Check error on rows.Close()
		defer func() {
			if err := rows.Close(); err != nil {
				log.Printf("failed to close rows: %v", err)
			}
		}()

		var students []Student
		for rows.Next() {
			var s Student
			if err := rows.Scan(&s.ID, &s.Name, &s.Age, &s.Email); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			students = append(students, s)
		}

		c.JSON(http.StatusOK, students)
	}
}
