package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Database {
    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
        log.Println("Warning: MONGODB_URI not set. Using default: mongodb://localhost:27017")
        uri = "mongodb://localhost:27017"
    }

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("❌ Failed to connect MongoDB:", err)
        return nil
    }

    dbName := os.Getenv("MONGODB_DB")
    if dbName == "" {
        log.Fatal("❌ MONGODB_DB not set in .env")
        return nil
    }

    log.Println("✅ Connected to MongoDB:", dbName)
    return client.Database(dbName)
}

