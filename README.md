# 🎓 Student CRUD REST API  
> Golang + Gin + PostgreSQL + Prometheus + Makefile

---

## 🧭 Overview

A production-grade, versioned RESTful API for managing student records — built with **Go** and the **Gin** web framework, backed by a **PostgreSQL** database. Designed with clean architecture, environment configuration, and monitoring readiness.

---

## 📌 Key Features

- ✅ CRUD operations on student records
- 🔐 Environment-based configuration via `.env`
- 🧪 Unit tests using `sqlmock`
- 📊 Prometheus-compatible `/metrics` endpoint
- 🔁 Healthcheck endpoint `/healthcheck`
- 🔧 Makefile automation (`build`, `run`, `test`, `migrate`)
- 📬 Postman collection for API testing

---

## 📖 What the API Does

The API supports:

- **Create** a new student  
- **Read** all students or a specific student by ID  
- **Update** existing student details  
- **Delete** a student record  

Additional endpoints:

- `/healthcheck` — Readiness status  
- `/metrics` — Prometheus-compatible metrics for observability

---

## 🗂️ Project Structure

```
One2N-SRE-Bootcamp/
├── main.go                 # Entry point: loads config, sets up DB & starts Gin server
├── config.go               # Loads environment variables from .env
├── db.go                   # Connects to PostgreSQL using environment config
├── model.go                # Defines Student struct for ORM mapping
├── handler.go              # Implements all HTTP handlers with logging
├── routes.go               # Registers routes and middleware with Gin engine
├── prometheus.go           # (Optional) Prometheus metrics middleware & endpoint
├── schema.sql              # SQL schema to create 'students' table
├── Dockerfile              # To containerize the application
├── .dockerignore           # Exclude unnecessary files 
├── .env                    # Environment variables (excluded from version control)
├── Makefile                # Build, run, migrate, test automation targets
├── handler_test.go         # Unit tests using sqlmock for isolated handler testing
├── postman_collection.json # Postman collection for testing all API endpoints
├── go.mod / go.sum         # Go module files for dependency tracking
```
---

## 🚀 Getting Started

### 1. 📦 Prerequisites

- [Go](https://golang.org/dl/) installed  
- [PostgreSQL](https://www.postgresql.org/download/) installed and running locally  
- [`psql`](https://www.postgresql.org/docs/current/app-psql.html) CLI access  

### 2. 🛠️ Setup PostgreSQL

Create a database and user:

```sql
CREATE USER studentuser WITH PASSWORD 'yourpassword';
CREATE DATABASE studentdb OWNER studentuser;
GRANT ALL PRIVILEGES ON DATABASE studentdb TO studentuser;
```

3. 🔐 **Create `.env` file**

Populate the `.env` (example below):

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=studentuser
DB_PASSWORD=yourpassword
DB_NAME=studentdb
DB_SSLMODE=disable
PORT=8080
```


4. 🔁 **Migrate database schema**

Use Makefile or run manually:

```
make migrate
```
### OR
```
psql -h localhost -U studentuser -d studentdb -f schema.sql
```


5. ▶️ **Build and run the server**

Use:
```
make run
```
### OR
```
go run .
```


6. 🧪 **Test the API**

- Use the included `postman_collection.json` imported into Postman.
- Or use `curl` commands for endpoints such as:

  ```
  # Health check
  curl http://localhost:8080/healthcheck
  ```
  ```
  # ➕ Add a student
  curl -X POST http://localhost:8080/api/v1/students \
    -H "Content-Type: application/json" \
    -d '{"name":"Alice","age":22,"email":"alice@example.com"}'
  ```

  ---

## 🧪 Testing

- Run tests locally:

```
make test
```
### OR
```
go test -v
```

---

## 📊 Monitoring

- A Prometheus-compatible `/metrics` endpoint is exposed.
- Middleware collects HTTP request metrics such as counts and durations.
- You can integrate this API with Prometheus & Grafana to visualize service metrics and health.

---
