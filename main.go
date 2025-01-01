package main

import (
	"log"
	"net/http"
	"os"

	"meta_ID_backend/database"
	"meta_ID_backend/resolvers"

	"github.com/graphql-go/handler"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB()

	schema := resolvers.NewRoot()

	h := handler.New(&handler.Config{
		Schema:   schema,
		GraphiQL: true,
	})

	// CORS 설정을 끄기 위해 직접 핸들러를 설정
	http.Handle("/graphql", h)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
