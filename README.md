# PROIS Backend API

PROIS Backend API is a RESTful service for [PROIS (Procurement Information System)](https://github.com/jonimaulanayusuf/prois) built with **Fiber** and **GORM**.

---

## ⚡ Postman Collection

Find the complete **Postman Collection** for testing and exploring all available API endpoints here:
- [PROIS Backend API – Postman Collection](https://github.com/user-attachments/files/24330782/Prois.Backend.postman_collection.json)

---

## ✅ Requirements

- Go 1.24+
- MySQL
- Git

---

## 🚀 Installation & Setup

### 1. Clone Repository

```bash
git clone https://github.com/jonimaulanayusuf/prois-backend.git
cd prois-backend
```

### 2. Environment Setup

Copy the environment example file:

```bash
cp .env.example .env
```

Configure the database, purchase webhook, JWT secret, and any other required settings in the .env file:

```bash
APP_ENV=local
APP_PORT=3001

DB_HOST=127.0.0.1
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=
DB_NAME=prois_db

JWT_SECRET= # example: verysecret
ENC_SECRET= # example: verysecret
PURCHASE_WEBHOOK_URL= # example: https://webhook.site/1fef1af8-2f72-4a28-b8e6-8ed7e88a5259
ALLOWED_ORIGINS= # default is "*", example: http://localhost:5500
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run The Server

```bash
go run cmd/main.go
```

Backend server will be running at:
```
http://localhost:3001
```
