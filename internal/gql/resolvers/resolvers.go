// Package resolvers contains GraphQL resolvers
package resolvers

import (
	"github.com/mortenoj/go-graphql-template/internal/gql"
	"github.com/mortenoj/go-graphql-template/internal/orm"
)

// Resolver is a modifable struct that can be used to pass on properties used
// in the resolvers, such as DB access
type Resolver struct {
	ORM *orm.ORM
}

// Mutation exposes mutation methods
func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

// Query exposes query methods
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
