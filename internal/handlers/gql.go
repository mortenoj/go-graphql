package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-gonic/gin"
	"github.com/mortenoj/go-graphql-template/internal/gql"
	"github.com/mortenoj/go-graphql-template/internal/gql/resolvers"
	"github.com/mortenoj/go-graphql-template/internal/orm"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	c := gql.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm,
		},
	}

	h := handler.NewDefaultServer(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("Go GraphQL Server", path)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
