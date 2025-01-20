package main

import (
	"connect-api/database"
	"connect-api/routes"
	"log"

	"github.com/gin-gonic/gin"
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

	// ตั้งค่า Router
	router := gin.Default()

	// ลงทะเบียน Routes
	routes.UserRoutes(router)

	// เริ่มต้น Server
	router.Run(":8080")
}
