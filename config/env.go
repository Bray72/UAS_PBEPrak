package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBDsn      string
	JWTSecret  string
	ApiKey     string
	MongoURI   string
	MongoDB    string
}

var cfg *Config

func LoadEnv() *Config {
	if cfg != nil {
		return cfg
	}

	// Coba load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found, using system environment variables")
	}

	// Load or set default values
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}

	dbDsn := os.Getenv("DB_DSN")
	if dbDsn == "" {
		dbDsn = "postgres://postgres:admin@localhost:5432/uas_db?sslmode=disable"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-min-32-characters-long-for-security"
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "12345"
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	mongoDB := os.Getenv("MONGO_DB")
	if mongoDB == "" {
		mongoDB = "uas_db"
	}

	cfg = &Config{
		AppPort:  appPort,
		DBDsn:    dbDsn,
		JWTSecret: jwtSecret,
		ApiKey:    apiKey,
		MongoURI:  mongoURI,
		MongoDB:   mongoDB,
	}

	return cfg
}
