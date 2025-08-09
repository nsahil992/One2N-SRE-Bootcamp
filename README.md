# 🎓 Student CRUD REST API  
> Golang + Gin + PostgreSQL + Prometheus + Makefile

---

## 🧭 Overview

A production-grade, versioned RESTful API for managing student records. Built with **Go** and the **Gin** web framework, backed by a **PostgreSQL** database. Designed with clean architecture, environment configuration, and monitoring readiness.

---

## 📌 Key Features

- ✅ CRUD operations on student records
- 🔐 Environment-based configuration via `.env`
- 🧪 Unit tests for different endpoints
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
│
├── .github/                      # GitHub-specific files
│   └── workflows/
│       └── cicd.yaml             # GitHub Actions CI/CD pipeline
│
├── argocd/                       # ArgoCD configuration
│   └── argocd-app.yml            # ArgoCD application manifest
│
├── helm/                         # Helm chart directory
│   ├── templates/
│   │   ├── application.yml       # Kubernetes Deployment
│   │   ├── database.yml          # Kubernetes DB Deployment
│   │   ├── postgres-pv.yml       # PersistentVolume for PostgreSQL
│   │   └── postgres-secret.yml   # Secrets for PostgreSQL
│   ├── Chart.yaml                # Helm chart definition
│   └── values.yaml               # Default values for the Helm chart
├── routes.go                     # Registers routes and middleware with Gin engine
├── schema.sql                    # SQL schema to create 'students' table
├── scripts/                      # Utility scripts
│   └── devtools.sh               # Tools installation script
├── prometheus.go                 # Prometheus metrics middleware & endpoint
├── student-data/                 # Directory to store student-related data
├── model.go                      # Defines Student struct 
├── .dockerignore                 # Exclude unnecessary files 
├── .env                          # Environment variables (excluded from version control)
├── .gitignore                    # Git ignore file
├── config.go                     # App configuration
├── db.go                         # Connects to PostgreSQL using environment config
├── docker-compose.yaml           # Docker Compose config for local setup
├── Dockerfile                    # Dockerfile for building the application image
├── go.mod / go.sum               # Go module files for dependency tracking
├── handler.go                    # Implements all HTTP handlers with logging
├── handler_test.go               # Unit tests using sqlmock for isolated handler testing
├── main.go                       # Application entry point
├── makefile                      # Makefile automation targets
└── migrate.go                    # Migration logic

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

## 🐳 Containerise REST API

### Instructions to build the image and run the docker container

```
# Building the Docker image:
make docker-build
OR
docker build -t yourdockerusername/student-api:tag .
```

```
# Running the Docker image:
make docker-run
OR
docker run --env-file .env -p 8080:8080 yourdockerusername/student-api:tag
```

```
# Tagging the docker image:
docker tag student-api:tag yourdockerusername/student-api:tag
```

```
# Pushing the docker image:
make docker-push
OR
docker push yourdockerusername/student-api:tag
```

---

## 📦📦 Setup one-click local development setup

```
# Pre-requisite for existing tools that must already be installed (Only for Linux)

cd scripts && chmod +x devtools.sh
./devtools.sh
```

```
# Run docker compose
make docker-compose
OR
docker-compose up # Run DB + API using docker-compose
```

---

## ☸️ Deploy REST API & its dependent services in K8s

```
# Note: The commands can change according to the file structure

cd templates

kubectl create namespace student-api (Create a namespace for k8s isolation)

kubectl apply -f postgres-secret.yml -f postgres-pv.yml -f database.yml -f application.yml -n student-api

```

---

## 📈 Deploy REST API & its dependent services using Helm Charts

