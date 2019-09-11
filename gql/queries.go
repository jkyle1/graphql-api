package gql

import (
	"../../graphql-api/postgres"
	"github.com/graphql-go/graphql"
)

//holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

//NewRoot returns base query type
func NewRoot(db *postgres.Db) *Root {
	//create resolver holding db
	resolver := Resolver{db: db}

	//create a new Root that describes base query setup
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						//slice of User type
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}
	return &root
}
