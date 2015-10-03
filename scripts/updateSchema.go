package main

import (
	"encoding/json"
	"github.com/chris-ramon/graphql-go"
	"github.com/chris-ramon/graphql-go/testutil"
	"github.com/chris-ramon/graphql-go/types"
	"github.com/sogko/golang-relay-starter-kit/data"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Save JSON of full schema introspection for Babel Relay Plugin to use
	resultChannel := make(chan *types.GraphQLResult)
	go gql.Graphql(gql.GraphqlParams{
		Schema:        data.Schema,
		RequestString: testutil.IntrospectionQuery,
	}, resultChannel)
	result := <-resultChannel
	if result.HasErrors() {
		log.Fatalf("ERROR introspecting schema: %v", result.Errors)
		return
	} else {
		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
		err = ioutil.WriteFile("../data/schema.json", b, os.ModePerm)
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}

	}
	// TODO: Save user readable type system shorthand of schema
	// pending implementation of printSchema
	/*
		fs.writeFileSync(
		  path.join(__dirname, '../data/schema.graphql'),
		  printSchema(Schema)
		);
	*/
}
