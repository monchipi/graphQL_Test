package resolvers

import (
    "database/sql"
    "github.com/monchipi/ex-myserver/graph/generated"
)

type Resolver struct {
    DB *sql.DB
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

//func (r *Resolver) Todo() generated.TodoResolver {return &todoResolver{r}}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
