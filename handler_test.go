package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestHealthCheck(t *testing.T) {
	r := gin.Default()
	r.GET("/healthcheck", HealthCheck)

	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	expectedBody := `{"status":"OK"}`
	if strings.TrimSpace(w.Body.String()) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}

func TestCreateStudent(t *testing.T) {
	db, mock, _ := sqlmock.New()
	r := gin.Default()
	r.POST("/api/v1/students", CreateStudent(db))

	mock.ExpectQuery(`INSERT INTO students`).
		WithArgs("Test User", 20, "test@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	body := `{"name":"Test User","age":20,"email":"test@example.com"}`
	req, _ := http.NewRequest("POST", "/api/v1/students", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", w.Code)
	}
}

func TestGetStudent(t *testing.T) {
	db, mock, _ := sqlmock.New()
	r := gin.Default()
	r.GET("/api/v1/students/:id", GetStudent(db))

	row := sqlmock.NewRows([]string{"id", "name", "age", "email"}).AddRow(1, "Test User", 20, "test@example.com")
	mock.ExpectQuery(`SELECT id, name, age, email FROM students WHERE id = \$1`).
		WithArgs("1").WillReturnRows(row)

	req, _ := http.NewRequest("GET", "/api/v1/students/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestUpdateStudent(t *testing.T) {
	db, mock, _ := sqlmock.New()
	r := gin.Default()
	r.PUT("/api/v1/students/:id", UpdateStudent(db))

	mock.ExpectExec(`UPDATE students`).
		WithArgs("Test User Updated", 21, "testupdated@example.com", "1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	body := `{"name":"Test User Updated","age":21,"email":"testupdated@example.com"}`
	req, _ := http.NewRequest("PUT", "/api/v1/students/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestDeleteStudent(t *testing.T) {
	db, mock, _ := sqlmock.New()
	r := gin.Default()
	r.DELETE("/api/v1/students/:id", DeleteStudent(db))

	mock.ExpectExec(`DELETE FROM students WHERE id = \$1`).
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	req, _ := http.NewRequest("DELETE", "/api/v1/students/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
