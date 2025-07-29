package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

// Helper function to create a Gin engine with routes and mocked DB
func setupRouterWithDB(mockDB sqlmock.Sqlmock, dbRows ...func()) (*gin.Engine, *sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("failed to create sqlmock: " + err.Error())
	}
	if mockDB != nil {
		mock = mockDB
	}
	r := gin.Default()

	// Attach Prometheus middleware if you have it (optional)
	// r.Use(PrometheusMiddleware())

	api := r.Group("/api/v1")
	{
		api.POST("/students", CreateStudent(db))
		api.GET("/students", GetAllStudents(db))
		api.GET("/students/:id", GetStudent(db))
		api.PUT("/students/:id", UpdateStudent(db))
		api.DELETE("/students/:id", DeleteStudent(db))
	}

	r.GET("/healthcheck", HealthCheck)
	r.GET("/metrics", MetricsHandler())

	return r, &mock
}

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

	// Expect INSERT query with given args
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
	// Optionally, check response body contains id, etc.
}

func TestGetAllStudents(t *testing.T) {
	db, mock, _ := sqlmock.New()
	r := gin.Default()
	r.GET("/api/v1/students", GetAllStudents(db))

	// Prepare mock rows
	rows := sqlmock.NewRows([]string{"id", "name", "age", "email"}).
		AddRow(1, "Alice", 22, "alice@example.com").
		AddRow(2, "Bob", 21, "bob@example.com")

	mock.ExpectQuery(`SELECT id, name, age, email FROM students`).WillReturnRows(rows)

	req, _ := http.NewRequest("GET", "/api/v1/students", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	// Optionally check JSON contains the students array
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

func TestMetricsEndpoint(t *testing.T) {
	r := gin.Default()
	r.GET("/metrics", MetricsHandler())

	req, _ := http.NewRequest("GET", "/metrics", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
