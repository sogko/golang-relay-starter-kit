package main

import (
	"net/http"
	"github.com/sogko/graphql-relay-go"

	"github.com/sogko/golang-relay-starter-kit/data"
)

func main() {

	// simplest relay-compliant graphql server HTTP handler
	h := gqlrelay.NewHandler(&gqlrelay.HandlerConfig{
		Schema: &data.Schema,
		Pretty: true,
	})

	// create graphql endpoint
	http.Handle("/graphql", h)

	// serve!
	http.ListenAndServe(":8080", nil)
}
