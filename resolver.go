//go:generate gorunpkg github.com/99designs/gqlgen

package golang_graphql

import (
	context "context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateArtist(ctx context.Context, name string, genre *Genre) (Artist, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteArtist(ctx context.Context, name string) (string, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Artists(ctx context.Context) ([]Artist, error) {
	panic("not implemented")
}
func (r *queryResolver) Artist(ctx context.Context, name string) (Artist, error) {
	panic("not implemented")
}
