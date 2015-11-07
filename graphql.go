package main

import (
	"github.com/graphql-go/handler"
	"github.com/sogko/golang-relay-starter-kit/data"
	"log"
	"net/http"
)

func main() {

	// simplest relay-compliant graphql server HTTP handler
	h := handler.New(&handler.Config{
		Schema: &data.Schema,
		Pretty: true,
	})

	// create graphql endpoint
	http.Handle("/graphql", h)

	// serve!
	port := ":8080"
	log.Printf(`GraphQL server starting up on http://localhost%v`, port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
