package resolvers

import (
	"meta_ID_backend/graph"

	"github.com/graphql-go/graphql"
)

// Query 및 Mutation을 설정하는 함수
func NewRoot() *graphql.Schema {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graph.QueryType,
		Mutation: graph.MutationType,
	})
	return &schema
}
