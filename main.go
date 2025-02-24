package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files from React build folder
	r.Static("/static", "./frontend/build/static")

	// Serve specific React build files
	r.StaticFile("/favicon.ico", "./frontend/build/favicon.ico")
	r.StaticFile("/manifest.json", "./frontend/build/manifest.json")
	r.StaticFile("/logo192.png", "./frontend/build/logo192.png")
	r.StaticFile("/logo512.png", "./frontend/build/logo512.png")
	r.StaticFile("/robots.txt", "./frontend/build/robots.txt")

	// API endpoint for chat
	r.POST("/chat", chatHandler)

	// Serve React `index.html` for all other routes
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	log.Println("✅ Server running on http://localhost:8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}

func chatHandler(c *gin.Context) {
	// Your chat logic here
	c.JSON(http.StatusOK, gin.H{"message": "Hello from backend!"})
}
