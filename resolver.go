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
	addedArtist, err := artistRepositoryInstance.Add(name, genre)
	if err != nil {
		return Artist{}, err
	}
	return addedArtist, nil
}
func (r *mutationResolver) DeleteArtist(ctx context.Context, name string) (string, error) {
	return name, artistRepositoryInstance.Delete(name)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Artists(ctx context.Context) ([]Artist, error) {
	return artistRepositoryInstance.FindAll()
}
func (r *queryResolver) Artist(ctx context.Context, name string) (Artist, error) {
	return artistRepositoryInstance.FindByName(name)
}
