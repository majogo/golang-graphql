//go:generate gorunpkg github.com/99designs/gqlgen

package golang_graphql

import (
	context "context"
)

type Resolver struct{}

func (r *Resolver) Artist() ArtistResolver {
	return &artistResolver{r}
}
func (r *Resolver) Disc() DiscResolver {
	return &discResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type artistResolver struct{ *Resolver }

func (r *artistResolver) OwnDiscs(ctx context.Context, obj *Artist) ([]Disc, error) {
	return discRepositoryInstance.FindByArtistName(obj.Name)
}

type discResolver struct{ *Resolver }

func (r *discResolver) ContainingTracks(ctx context.Context, obj *Disc) ([]Track, error) {
	panic("not implemented")
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

func (r *queryResolver) Discs(ctx context.Context) ([]Disc, error) {
	return discRepositoryInstance.FindAll()
}
func (r *queryResolver) Disc(ctx context.Context, name string) (Disc, error) {
	return discRepositoryInstance.FindByName(name)
}

func (r *queryResolver) Tracks(ctx context.Context) ([]Track, error) {
	return trackRepositoryInstance.FindAll()
}

func (r *queryResolver) Track(ctx context.Context, name string) (Track, error) {
	return trackRepositoryInstance.FindByName(name)
}
