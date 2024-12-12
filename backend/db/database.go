package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDatabase() {
	// 加載 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default configuration")
	}

	// 從環境變數中讀取 MongoDB 連接 URI
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalf("MONGODB_URI is not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(uri)

	// 連接到 MongoDB
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	DB = client
	log.Println("Connected to MongoDB successfully")
}
