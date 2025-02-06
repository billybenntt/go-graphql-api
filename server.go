package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
	"go-graph-api/db"
	"go-graph-api/graph"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"
	// Hardcoded Neo4j credentials (modify as needed)
	neo4jURI      = "neo4j://100.100.20.30:7687"
	neo4jUser     = "neo4j"
	neo4jPassword = "abc123xxx"
)

func main() {

	// Initialize Neo4j connection with hardcoded credentials
	err := db.InitializeNeo4j(neo4jURI, neo4jUser, neo4jPassword)
	if err != nil {
		log.Fatalf("Database initialization error: %v", err)
	}
	defer db.CloseNeo4j() // Ensure connection closes on exit

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize GraphQL
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL Subscriptions", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 connect to http://localhost:%s/ for GraphQL Subscriptions", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
