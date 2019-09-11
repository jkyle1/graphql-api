package server

import (
	"encoding/json"
	"net/http"

	"../../graphql-api/gql"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

type Server struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

//GraphQL returns http.HandlerFunc for /graphql endpoint
func (s *Server) GraphQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check to ensure query was provided in request body
		if r.Body == nil {
			http.Error(w, "GraphQL query must be provided in query body", 400)
			return
		}

		var rBody reqBody
		//Decode request body into rBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		//execute GraphQL query
		result := gql.ExecuteQuery(rBody.Query, *s.GqlSchema)
		//handles marshalling to json, automatically escaping HTML and setting Content-Type to application/json
		render.JSON(w, r, result)
	}
}
