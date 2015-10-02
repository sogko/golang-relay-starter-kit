package data
import (
	"github.com/sogko/graphql-relay-go"
	"github.com/chris-ramon/graphql-go/types"
)

var userType *types.GraphQLObjectType
var widgetType *types.GraphQLObjectType

var nodeDefinitions *gqlrelay.NodeDefinitions
var widgetConnection *gqlrelay.GraphQLConnectionDefinitions

var Schema types.GraphQLSchema

func init() {

	/**
	 * We get the node interface and field from the Relay library.
	 *
	 * The first method defines the way we resolve an ID to its object.
	 * The second defines the way we resolve an object to its GraphQL type.
	 */
	nodeDefinitions = gqlrelay.NewNodeDefinitions(gqlrelay.NodeDefinitionsConfig{
		IdFetcher:  func(id string, info types.GraphQLResolveInfo) interface{} {
			resolvedId := gqlrelay.FromGlobalId(id)
			if resolvedId.Type == "User" {
				return GetUser(resolvedId.Id)
			}
			if resolvedId.Type == "Widget" {
				return GetWidget(resolvedId.Id)
			}
			return nil
		},
		TypeResolve: func(value interface{}, info types.GraphQLResolveInfo) *types.GraphQLObjectType {
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
	widgetType = types.NewGraphQLObjectType(types.GraphQLObjectTypeConfig{
		Name: "Widget",
		Description: "A shiny widget'",
		Fields: types.GraphQLFieldConfigMap{
			"id": gqlrelay.GlobalIdField("Widget", nil),
			"name": &types.GraphQLFieldConfig{
				Description: "The name of the widget",
				Type: types.GraphQLString,
			},
		},
		Interfaces: []*types.GraphQLInterfaceType{
			nodeDefinitions.NodeInterface,
		},
	})
	widgetConnection = gqlrelay.ConnectionDefinitions(gqlrelay.ConnectionConfig{
		Name: "WidgetConnection",
		NodeType: widgetType,
	})

	userType = types.NewGraphQLObjectType(types.GraphQLObjectTypeConfig{
		Name: "User",
		Description: "A person who uses our app",
		Fields: types.GraphQLFieldConfigMap{
			"id": gqlrelay.GlobalIdField("User", nil),
			"widgets": &types.GraphQLFieldConfig{
				Type: widgetConnection.ConnectionType,
				Description: "A person's collection of widgets",
				Args: gqlrelay.ConnectionArgs,
				Resolve: func(p types.GQLFRParams) interface{} {
					args := gqlrelay.NewConnectionArguments(p.Args)
					dataSlice := WidgetsToInterfaceSlice(GetWidgets()...)
					return gqlrelay.ConnectionFromArray(dataSlice, args)
				},
			},
		},
		Interfaces: []*types.GraphQLInterfaceType{
			nodeDefinitions.NodeInterface,
		},
	})

	/**
	 * This is the type that will be the root of our query,
	 * and the entry point into our schema.
	 */
	queryType := types.NewGraphQLObjectType(types.GraphQLObjectTypeConfig{
		Name: "Query",
		Fields: types.GraphQLFieldConfigMap{
			"node": nodeDefinitions.NodeField,

			// Add you own root fields here
			"viewer": &types.GraphQLFieldConfig{
				Type: userType,
				Resolve: func(p types.GQLFRParams) interface{} {
					return GetViewer()
				},
			},
		},
	})

	/**
	 * This is the type that will be the root of our mutations,
	 * and the entry point into performing writes in our schema.
	 */
	//	mutationType := types.NewGraphQLObjectType(types.GraphQLObjectTypeConfig{
	//		Name: "Mutation",
	//		Fields: types.GraphQLFieldConfigMap{
	//			// Add you own mutations here
	//		},
	//	})

	/**
	* Finally, we construct our schema (whose starting query type is the query
	* type we defined above) and export it.
	*/
	var err error
	Schema, err = types.NewGraphQLSchema(types.GraphQLSchemaConfig{
		Query: queryType,
	})
	if err != nil {
		panic(err)
	}

}