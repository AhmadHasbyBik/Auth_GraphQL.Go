package main

import (
	"graphql/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

//const defaultPort = "8080"

func main() {
	router := chi.NewRouter()
	router.Use(graph.Middleware())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	graph.Database()
	port := os.Getenv("PORT")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
