# PROIS Backend API

PROIS Backend API is a RESTful service for [PROIS (Procurement Information System)](https://github.com/jonimaulanayusuf/prois) built with **Fiber** and **GORM**.

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

Configure the database, purchase webhook, JWT secret, and any other required settings in the .env file.

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