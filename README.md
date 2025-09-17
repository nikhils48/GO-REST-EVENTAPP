# Gin REST API with SQLite Migrations

This repository contains a **RESTful API built with Go and Gin**, using **SQLite** as the database and **golang-migrate** for schema migrations.  
It is designed for **learning, portfolio, and demo purposes**.

---

## Features

- REST API endpoints using **Gin framework**
- Database migrations using **golang-migrate**
- SQLite database (lightweight, file-based)
- Easy setup for local development
- Ready to extend for additional features

---

## Prerequisites

- [Go](https://golang.org/dl/) 1.20+ installed
- Git installed
- Optional: `modernc.org/sqlite` driver installed via `go get`

---

## Setup

1. **Clone the repository**

`bash`
git clone 

2. Run migrations to create the database

go run ./cmd/migrate/main.go up


This will create data/data.db (SQLite database) and apply all migrations.

Make sure the data/ folder exists (it will be created automatically if missing).

3. Start the server

go run main.go


The API will be available at http://localhost:8080 by default.
