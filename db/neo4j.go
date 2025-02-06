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

// InitializeNeo4j sets up the Neo4j connection
func InitializeNeo4j(uri, username, password string) error {
	var err error
	once.Do(func() {
		driver, err = neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
		if err != nil {
			log.Fatalf("Failed to connect to Neo4j: %v", err)
		}
	})

	// Verify the connection
	ctx := context.Background()
	if err = driver.VerifyConnectivity(ctx); err != nil {
		return fmt.Errorf("neo4j connectivity error: %v", err)
	}

	log.Println("Connected to Neo4j successfully")
	return nil
}

// GetDriver returns the Neo4j driver instance
func GetDriver() neo4j.DriverWithContext {
	return driver
}

// CloseNeo4j closes the database connection
func CloseNeo4j() {
	if driver != nil {
		_ = driver.Close(context.Background())
		log.Println("Closed Neo4j connection")
	}
}
