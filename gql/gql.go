package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

//ExecuteQuery runs GraphQL queries
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("Errors in ExecuteQuery: %v", result.Errors)
	}
	return result
}
