package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/htanmo/hackernews/graph"
	"github.com/htanmo/hackernews/internal/auth"
	"github.com/htanmo/hackernews/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.ConnectDB()
	defer database.CloseDB()

	router := http.NewServeMux()

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	authMiddleware := auth.Middleware()
	router.Handle("/query", authMiddleware(srv))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Printf("API endpoint is available at http://localhost:%s/query", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
