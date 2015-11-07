# Golang-Relay Starter Kit

This kit includes:
- a **NodeJS** app server: to serve the front-end written with `react-relay`
- a **Golang** GraphQL server: to serve the back-end `graphql-go` server that handles GraphQL queries
- a Babel transpiler workflow using `webpack` that you can use to get started building an app with Relay.

For a walkthrough, see the [Relay tutorial](https://facebook.github.io/relay/docs/tutorial.html).

### Notes:
This is based on alpha version of `graphql-go` and `graphql-relay-go`. 
Be sure to watch both repositories for latest changes.

## Installation

- Install dependencies for NodeJS app server
```
npm install
```
- Install dependencies for Golang GraphQL server
```
go get -v
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

### JavaScript
Any changes you make to files in the `js/` directory will cause the server to
automatically rebuild the app and refresh your browser.

### Golang

#### Schema data
Since Golang does not support loading package / module dynamically, remember to update the package import for schema data in:
- `graphql.go`
- `scripts/updateSchema.go`

For e.g

```go
import (
  ...
  "github.com/sogko/golang-relay-starter-kit/data" // <--- update to package containing schema
)
```

#### Schema updates
If at any time you make changes to `data/schema.go`, stop the server,
regenerate `data/schema.json`, and restart the server:

```
npm run update-schema
npm start
```

`schema.json` is needed by the JS code for `./build/babelRelayPlugin.js`

## Examples
- [todomvc-relay-go](https://github.com/sogko/todomvc-relay-go) - Port of the React/Relay TodoMVC app, driven by a Golang GraphQL backend

Feel free to submit a PR to add to this list.

## TODOs
- [x] Swap out `express-graphql` server with a Golang GraphQL server
- [x] GraphQL schema definition in Golang
- [x] Generate `schema.json` from schema definition for `babel-relay-plugin`
- [ ] Generate `schema.graphql` from schema definition

## Credits
This kit is build on top of https://github.com/relayjs/relay-starter-kit