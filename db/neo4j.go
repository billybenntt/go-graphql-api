package db

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"sync"
)

// Global variables for connection pooling
var (
	driver neo4j.DriverWithContext
	once   sync.Once
)

const (
	// Hardcoded Neo4j credentials (modify as needed)
	neo4jURI      = "neo4j://100.100.20.30:7687"
	neo4jUser     = "neo4j"
	neo4jPassword = "abc123xxx"
)

// InitDB sets up the Neo4j connection
func InitDB() error {
	var err error
	once.Do(func() {
		driver, err = neo4j.NewDriverWithContext(neo4jURI, neo4j.BasicAuth(neo4jUser, neo4jPassword, ""))
		if err != nil {
			log.Fatalf("Failed to connect to Neo4j: %v", err)
		}
	})

	// Verify the connection
	ctx := context.Background()

	// error handling
	if err = driver.VerifyConnectivity(ctx); err != nil {
		return fmt.Errorf("neo4j connectivity error: %v", err)
	}

	log.Println("Connected to Neo4j successfully")
	return nil
}

// GetDBConnection returns the Neo4j driver instance
func GetDBConnection() neo4j.DriverWithContext {
	return driver
}

// CloseDB closes the database connection
func CloseDB() {
	if driver != nil {
		_ = driver.Close(context.Background())
		log.Println("Closed Neo4j connection")
	}
}
