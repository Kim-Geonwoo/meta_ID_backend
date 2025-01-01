package graph

import (
	"meta_ID_backend/database"
	"meta_ID_backend/models"

	"github.com/graphql-go/graphql"
)

// UserType 정의
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// QueryType 정의
var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(UserType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var users []models.User
					database.DB.Find(&users)
					return users, nil
				},
			},
			"user": &graphql.Field{
				Type: UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					var user models.User
					database.DB.First(&user, id)
					return user, nil
				},
			},
		},
	},
)

// MutationType 정의
var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: UserType,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user := models.User{
						Username: p.Args["username"].(string),
						Email:    p.Args["email"].(string),
						Password: p.Args["password"].(string),
					}
					database.DB.Create(&user)
					return user, nil
				},
			},
			"updateUser": &graphql.Field{
				Type: UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					var user models.User
					database.DB.First(&user, id)

					if username, ok := p.Args["username"].(string); ok {
						user.Username = username
					}
					if email, ok := p.Args["email"].(string); ok {
						user.Email = email
					}
					if password, ok := p.Args["password"].(string); ok {
						user.Password = password
					}

					database.DB.Save(&user)
					return user, nil
				},
			},
			"deleteUser": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, _ := p.Args["id"].(string)
					database.DB.Delete(&models.User{}, id)
					return true, nil
				},
			},
		},
	},
)

// GraphQL 스키마 정의
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	},
)
