package main

import (
	"log"
	"net/http"
	"os"

	"meta_ID_backend/database"
	"meta_ID_backend/resolvers"

	"github.com/rs/cors"

	"github.com/graphql-go/handler"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// 데이터베이스 초기화
	database.InitDB()

	// GraphQL 스키마 설정
	schema := resolvers.NewRoot()

	// GraphQL 핸들러 설정
	h := handler.New(&handler.Config{
		Schema:   schema,
		GraphiQL: true,
	})

	// CORS 설정 추가
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"metaid-backend.geonwoo.dev"}, // 허용할 클라이언트의 도메인
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	}).Handler(h)

	http.Handle("/graphql", corsHandler)

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
