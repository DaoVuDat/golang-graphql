package main

import (
	"database/sql"
	"github.com/DaoVuDat/graphql/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func setupGraphqlHandler(db *sql.DB) (*handler.Server, http.HandlerFunc) {
	graphqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{
				Db: db,
			}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL playground", "/graphql")
	return graphqlHandler, playgroundHandler
}
