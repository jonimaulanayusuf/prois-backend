package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"prois-backend/internal/config"
	"prois-backend/internal/models"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("DB_USER", "root"),
		config.GetEnv("DB_PASSWORD", ""),
		config.GetEnv("DB_HOST", "127.0.0.1"),
		config.GetEnv("DB_PORT", "3306"),
		config.GetEnv("DB_NAME", "prois_db"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Supplier{},
		&models.Purchasing{},
		&models.PurchasingDetail{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	} else {
		log.Println("Table user auto-migrated")
	}

	DB = db
	log.Println("Database connected")
}
