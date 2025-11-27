package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    // Coba load file .env
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️  .env file not found, using system environment variables")
    }

    // Default values jika belum ada di environment
    if os.Getenv("APP_PORT") == "" {
        os.Setenv("APP_PORT", "3000")
    }

    // Tambahkan untuk MongoDB
    if os.Getenv("MONGO_URI") == "" {
        os.Setenv("MONGO_URI", "mongodb://localhost:27017")
    }

    if os.Getenv("MONGO_DB") == "" {
        os.Setenv("MONGO_DB", "mhs_db")
    }
}
