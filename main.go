package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"laundry-api/config"
	"laundry-api/routes"
)

func main() {
	// 1️⃣ Load file .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Tidak menemukan file .env, lanjut dengan environment default")
	}

	// 2️⃣ Inisialisasi koneksi database
	config.InitDB()

	// 3️⃣ Buat instance router Gin
	r := gin.Default()

	// 4️⃣ Middleware CORS (🔥 ini versi lengkap dan aman)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 5️⃣ Tambahkan handler global untuk preflight OPTIONS
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	// 6️⃣ Setup semua route (setelah CORS)
	routes.SetupRoutes(r)

	// 7️⃣ Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("✅ Backend running at http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Gagal menjalankan server: %v", err)
	}
}
