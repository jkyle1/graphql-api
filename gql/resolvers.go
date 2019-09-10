package gql

import (
	"../../graphql-api/postgres/postgres.go"
	"github.com/graphql-go/graphql"
)

//Resolver struct holds connection to db
type Resolver struct {
	db *postgres.Db
}

//UserResolver resolves our user query through a call to GetUserByName
func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	//strip name from args and assert is a string
	name, ok := p.Args["name"].(string)
	if ok {
		users := r.db.GetUsersByName(name)
		return users, nil
	}
	return nil, nil
}
