# Golang-Relay Starter Kit

This kit includes a **NodeJS** app server, a **Golang** GraphQL server, and a transpiler that you can use to get started building an app with Relay. For a walkthrough, see the [Relay tutorial](https://facebook.github.io/relay/docs/tutorial.html).

## Important!
This kit is build on top of https://github.com/relayjs/relay-starter-kit with the following goal:
- To replace `express-graphql` server with a working Golang GraphQL server using [`graphql-go`](https://github.com/chris-ramon/graphql-go) and [`graphq-relay-go`](https://github.com/sogko/graphql-relay-go)

Notes:
- This is based on alpha version of `graphql-go` and `graphql-relay-go`. Be sure to watch both repositories for changes.

## Installation


1. Install dependencies for NodeJS app server
```
npm install
```

2. Install dependencies for Golang GraphQL server
```
go get -v ./...
```

## Running

Start a local server:

```
npm start
```

The above command will run both the NodeJS app server and Golang GraphQL server concurrently.

- Golang GraphQL server will be running at http://localhost:8080/graphql
- NodeJS app server will be running at http://localhost:3000

## Developing

Any changes you make to files in the `js/` directory will cause the server to
automatically rebuild the app and refresh your browser.

If at any time you make changes to `data/schema.go`, stop the server,
regenerate `data/schema.json`, and restart the server:

```
npm run update-schema
npm start
```

## TODOs
- [x] Swap out `express-graphql` server with a Golang GraphQL server
- [x] GraphQL schema definition in Golang
- [x] Generate `schema.json` from schema definition for `babel-relay-plugin`
- [ ] Generate `schema.graphql` from schema definition


## License

Relay Starter Kit is [BSD licensed](./LICENSE). We also provide an additional [patent grant](./PATENTS).
