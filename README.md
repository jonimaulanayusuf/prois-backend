# PROIS Backend API

PROIS Backend API is a RESTful service for PROIS (Procurement Information System) built with **Golang** using a modular and scalable architecture.

This project is designed to provide a clean, maintainable backend foundation for modern web applications.

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
cd prois
```

### 2. Environment Setup

Copy the environment example file:

```bash
cp .env.example .env
vim .env
```

Configure the database and application settings in the .env file.

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
http://localhost:3000
```