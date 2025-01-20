package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client *mongo.Client
	Ctx    = context.TODO()
)

// ConnectDB เชื่อมต่อ MongoDB และคืนค่า client
func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // เปลี่ยน URI ตามต้องการ
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// ทดสอบการเชื่อมต่อ
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	Client = client
	return client
}

// GetCollection คืนค่าคอลเลกชันตามชื่อ
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return Client.Database(databaseName).Collection(collectionName)
}
