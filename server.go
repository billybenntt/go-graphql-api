package main

import (
	"github.com/gin-gonic/gin"
	"go-graph-api/db"
	"go-graph-api/routes"
	"log"
)

func main() {

	// Initialize Neo4j connection
	err := db.InitDB()

	if err != nil {
		log.Fatalf("Database initialization error: %v", err)
	}
	// close on exit
	defer db.CloseDB()

	// Start Server
	server := gin.Default()
	// Register all Routes with the Server
	routes.RegisterRoutes(server)
	log.Printf("🚀 connect to http://localhost:%d/api for GraphQL API", 8010)
	log.Printf("🚀 connect to http://localhost:%d/play for GraphQL Playground", 8010)
	server.Run(":8010")

}
