package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 加載 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 從環境變數中讀取 MongoDB URI
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatalf("MONGODB_URI is not set in .env file")
	}

	// 設置上下文和超時時間
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 使用 mongo.Connect 建立並連接客戶端
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// 延遲執行斷開連接
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	fmt.Println("Connected to MongoDB successfully!")

	// 測試列出資料庫名稱
	databases, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		log.Fatalf("Failed to list databases: %v", err)
	}
	fmt.Println("Databases:", databases)

	// 測試 sample_mflix 資料庫是否存在
	testSampleMflix(ctx, client)

	// 關閉連接
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}

	fmt.Println("Disconnected from MongoDB.")
}

// 測試 sample_mflix 資料庫中的集合和文檔
func testSampleMflix(ctx context.Context, client *mongo.Client) {
	fmt.Println("\nTesting sample_mflix database...")

	// 指定 sample_mflix 資料庫和 movies 集合
	moviesCollection := client.Database("sample_mflix").Collection("movies")

	// 檢查集合中文檔數量
	count, err := moviesCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to count documents in sample_mflix.movies: %v", err)
	}
	fmt.Printf("sample_mflix.movies contains %d documents\n", count)

	if count > 0 {
		fmt.Println("\nFetching some sample documents from sample_mflix.movies...")

		cursor, err := moviesCollection.Find(ctx, bson.M{}, options.Find().SetLimit(5))
		if err != nil {
			log.Fatalf("Failed to fetch documents from sample_mflix.movies: %v", err)
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var doc bson.M
			if err := cursor.Decode(&doc); err != nil {
				log.Fatalf("Failed to decode document: %v", err)
			}
			fmt.Println(doc) // 打印每個文檔內容
		}

		if err := cursor.Err(); err != nil {
			log.Fatalf("Cursor encountered an error: %v", err)
		}
	} else {
		fmt.Println("No documents found in sample_mflix.movies.")
	}
}
