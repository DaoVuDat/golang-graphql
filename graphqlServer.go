package main

import (
	"database/sql"
	"github.com/DaoVuDat/graphql/graph"
	"github.com/DaoVuDat/graphql/sqlStore"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func setupGraphqlHandler(db *sql.DB) (*handler.Server, http.HandlerFunc) {
	// Create Store
	store := sqlStore.NewStore(db)

	var graphqlHandler = handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{
				Store: store,
			}},
		),
	)

	//graphqlHandler = sqlStore.Middleware(db, graphqlHandler)

	playgroundHandler := playground.Handler("GraphQL playground", "/graphql")
	return graphqlHandler, playgroundHandler
}
