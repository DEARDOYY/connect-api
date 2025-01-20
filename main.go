package main

import (
	"connect-api/database"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// MongoDB client
var client *mongo.Client

func main() {
	// เชื่อมต่อ MongoDB
	client := database.ConnectDB()
	defer func() {
		if err := client.Disconnect(database.Ctx); err != nil {
			log.Fatalf("Error disconnecting MongoDB: %v", err)
		}
	}()
}
