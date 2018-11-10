package golang_graphql

import "fmt"

type artistRepository struct {
	artists []Artist
}

var artistRepositoryInstance artistRepository

func init() {
	artistRepositoryInstance = artistRepository{
		artists: make([]Artist, 0),
	}
}

func (repository *artistRepository) Add(name string, genre *Genre) (Artist, error) {
	newArtist := Artist{Name: name}
	if genre != nil {
		newArtist.Genre = genre
	}
	repository.artists = append(repository.artists, newArtist)
	return newArtist, nil
}

func (repository *artistRepository) Delete(name string) error {
	for i, artist := range repository.artists {
		if name == artist.Name {
			repository.artists = append(repository.artists[:i], repository.artists[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("artist with name name='%s' not found", name)
}

func (repository *artistRepository) FindAll() ([]Artist, error) {
	return repository.artists, nil
}

func (repository *artistRepository) FindByName(name string) (Artist, error) {
	for _, artist := range repository.artists {
		if artist.Name == name {
			return artist, nil
		}
	}
	return Artist{}, fmt.Errorf("artist with name='%v' not found", name)
}
