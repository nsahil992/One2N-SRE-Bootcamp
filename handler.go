package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// Basic example, metrics endpoint just returns up for now
func Metrics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"metrics": "up"})
}

func CreateStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var s Student
		if err := c.ShouldBindJSON(&s); err != nil {
			log.Printf("[WARN] Bad request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := db.QueryRow("INSERT INTO students (name, age, email) VALUES ($1, $2, $3) RETURNING id",
			s.Name, s.Age, s.Email).Scan(&s.ID)
		if err != nil {
			log.Printf("[ERROR] Inserting student: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Printf("[INFO] Created student with ID %d", s.ID)
		c.JSON(http.StatusCreated, s)
	}
}

func GetAllStudents(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, age, email FROM students")
		if err != nil {
			log.Printf("[ERROR] Fetching students: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()
		var students []Student
		for rows.Next() {
			var s Student
			if err := rows.Scan(&s.ID, &s.Name, &s.Age, &s.Email); err != nil {
				log.Printf("[ERROR] Scanning student: %v", err)
				continue
			}
			students = append(students, s)
		}
		log.Printf("[INFO] Fetched %d students", len(students))
		c.JSON(http.StatusOK, students)
	}
}

func GetStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var s Student
		err := db.QueryRow("SELECT id, name, age, email FROM students WHERE id = $1", id).
			Scan(&s.ID, &s.Name, &s.Age, &s.Email)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		} else if err != nil {
			log.Printf("[ERROR] Fetching student: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Printf("[INFO] Fetched student with ID %s", id)
		c.JSON(http.StatusOK, s)
	}
}

func UpdateStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var s Student
		if err := c.ShouldBindJSON(&s); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := db.Exec("UPDATE students SET name = $1, age = $2, email = $3 WHERE id = $4",
			s.Name, s.Age, s.Email, id)
		if err != nil {
			log.Printf("[ERROR] Updating student: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		n, _ := res.RowsAffected()
		if n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		log.Printf("[INFO] Updated student with ID %s", id)
		c.JSON(http.StatusOK, gin.H{"message": "Student updated"})
	}
}

func DeleteStudent(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		res, err := db.Exec("DELETE FROM students WHERE id = $1", id)
		if err != nil {
			log.Printf("[ERROR] Deleting student: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		n, _ := res.RowsAffected()
		if n == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		log.Printf("[INFO] Deleted student with ID %s", id)
		c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
	}
}
