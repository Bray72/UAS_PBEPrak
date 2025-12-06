package main

import (
	"clean-arch/config"
	"clean-arch/database"
	"log"
 	_ "github.com/lib/pq"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "clean-arch/docs"
)

func main() {
	env := config.LoadEnv()
	db := database.ConnectDB()
	defer db.Close()

	app := fiber.New(fiber.Config{
		AppName: "UAS - Sistem Pelaporan Prestasi Mahasiswa",
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	port := env.AppPort
	log.Printf("🚀 Server starting on http://localhost:%s\n", port)
	log.Printf("📖 Swagger docs available at http://localhost:%s/swagger/index.html\n", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}

// package main

// import (
// 	"clean-arch/app/handler"
// 	"clean-arch/app/repository"
// 	"clean-arch/app/service"
// 	"clean-arch/config"
// 	"clean-arch/database"
// 	"clean-arch/route"
// 	"log"
// 	"github.com/joho/godotenv"
// 	_ "github.com/lib/pq"
// 	"fmt"
// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// 	_ "clean-arch/docs"
// )

// func main() {
// 	password := "123456"
//     hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//     fmt.Println(string(hash))

// 	if err := godotenv.Load(); err != nil {
//         log.Fatal("Error loading .env file")
//     }

// 	env := config.LoadEnv()
// 	db := database.ConnectDB()
// 	if db == nil {
//     	log.Fatal("❌ ConnectDB returned NIL")
// 	}

// 	defer db.Close()

// 	app := fiber.New(fiber.Config{
// 		AppName: "UAS - Sistem Pelaporan Prestasi Mahasiswa",
// 	})

// 	app.Use(logger.New())
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "*",
// 		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
// 		AllowHeaders: "Content-Type,Authorization",
// 	}))

// 	authRepo := repository.NewAuthRepository(db)
// 	authService := service.NewAuthService(authRepo)
// 	authHandler := handler.NewAuthHandler(authService)

// 	routes.SetupRoutes(app, authHandler)

// 	port := env.AppPort
// 	log.Printf("🚀 Server starting on http://localhost:%s\n", port)
// 	log.Printf("📖 Swagger docs available at http://localhost:%s/swagger/index.html\n", port)

// 	if err := app.Listen(":" + port); err != nil {
// 		log.Fatalf("❌ Server failed to start: %v", err)
// 	}

	
// }
