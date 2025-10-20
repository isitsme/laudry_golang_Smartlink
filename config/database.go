package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"laundry-api/models" // ✅ tambahkan ini agar bisa akses struct User, Service, Order
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// ✅ tambahkan bagian ini untuk migrasi tabel otomatis
	err = db.AutoMigrate(&models.User{}, &models.Service{}, &models.Order{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db
	log.Println("✅ Connected to PostgreSQL & Migrated models")
}
