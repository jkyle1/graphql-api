package main

import (
	"fmt"
	"github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"../graphql-api/gql/gql.go"
	"../graphql-api/postgres/postgres.go"
	"../graphql-api/server/server.go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"


)

func main() {
	//initialise our api & return a pointer to our router for http.ListenAndServe
	//and a pointer to our db to defer closing when main() is finished
	router, db := initialiseAPI()
	defer db.Close()

	//listen on port 4000 & log error then exit
	log.Fatal(http.ListenAndServe(":4000", router))
}

func initialiseAPI() (*chi.Mux, *postgres.Db) {
	//create a new router
	router := chi.NewRouter()

	//create a new connection to our pg db
	db, err := postgres.New(postgres.ConnString("localhost", 5432, "bradford", "go_graphql_db"), )

	if err != nil {
		log.Fatal(err)
	}

	//create our root query for graphql
	rootQuery := gql.NewRoot(db)
	//create new graphql schema passing in the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	//create a server struct to hold a pointer to db & address of graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	//add middleware to router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), //set content-type headers
		middleware.Logger, //log api request calls
		middleware.DefaultCompress, //compress results gzipping assets/json
		middleware.StripSlashes, //match paths with trailing slash then strip the slash and continue
		middleware.Recoverer, //recover from panics without server crash
		 )

	//create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db
}