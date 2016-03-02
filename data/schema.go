package data

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

var userType *graphql.Object
var widgetType *graphql.Object

var nodeDefinitions *relay.NodeDefinitions
var widgetConnection *relay.GraphQLConnectionDefinitions

var Schema graphql.Schema

func init() {

	/**
	 * We get the node interface and field from the Relay library.
	 *
	 * The first method defines the way we resolve an ID to its object.
	 * The second defines the way we resolve an object to its GraphQL type.
	 */
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ct context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			if resolvedID.Type == "User" {
				return GetUser(resolvedID.ID), nil
			}
			if resolvedID.Type == "Widget" {
				return GetWidget(resolvedID.ID), nil
			}
			return nil, nil
		},
		TypeResolve: func(value interface{}, info graphql.ResolveInfo) *graphql.Object {
			switch value.(type) {
			case *User:
				return userType
			case *Widget:
				return widgetType
			}
			return nil
		},
	})

	/**
	 * Define your own types here
	 */
	widgetType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Widget",
		Description: "A shiny widget'",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Widget", nil),
			"name": &graphql.Field{
				Description: "The name of the widget",
				Type:        graphql.String,
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})
	widgetConnection = relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "WidgetConnection",
		NodeType: widgetType,
	})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A person who uses our app",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"widgets": &graphql.Field{
				Type:        widgetConnection.ConnectionType,
				Description: "A person's collection of widgets",
				Args:        relay.ConnectionArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					args := relay.NewConnectionArguments(p.Args)
					dataSlice := WidgetsToInterfaceSlice(GetWidgets()...)
					return relay.ConnectionFromArray(dataSlice, args), nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	/**
	 * This is the type that will be the root of our query,
	 * and the entry point into our schema.
	 */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": nodeDefinitions.NodeField,

			// Add you own root fields here
			"viewer": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetViewer(), nil
				},
			},
		},
	})

	/**
	 * This is the type that will be the root of our mutations,
	 * and the entry point into performing writes in our schema.
	 */
	//	mutationType := graphql.NewObject(graphql.ObjectConfig{
	//		Name: "Mutation",
	//		Fields: graphql.Fields{
	//			// Add you own mutations here
	//		},
	//	})

	/**
	* Finally, we construct our schema (whose starting query type is the query
	* type we defined above) and export it.
	 */
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		panic(err)
	}
}
