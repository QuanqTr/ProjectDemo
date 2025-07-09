package main

import (
	"log"
	"os"
	"project-backend/internal/database"
	"project-backend/internal/handlers"
	"project-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	database.Connect()

	// Auto migrate models
	database.DB.AutoMigrate(&models.Student{})

	// Initialize Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API routes
	api := r.Group("/api/v1")
	{
		// Student routes
		api.GET("/students", handlers.GetStudents)
		api.POST("/students", handlers.CreateStudent)
		api.GET("/students/:id", handlers.GetStudent)
		api.PUT("/students/:id", handlers.UpdateStudent)
		api.DELETE("/students/:id", handlers.DeleteStudent)
		api.GET("/students/major", handlers.GetStudentsByMajor)
		api.GET("/students/status", handlers.GetStudentsByStatus)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
