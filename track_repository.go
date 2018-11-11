package golang_graphql

import "fmt"

type trackRepository struct {
	tracks []Track
}

var trackRepositoryInstance trackRepository

func init() {
	trackRepositoryInstance = trackRepository{
		tracks: make([]Track, 0),
	}
}

func (repository *trackRepository) Add(name string, discName, artistName string, duration *int) (Track, error) {
	newTrack := Track{Name: name, DiscName: discName, ArtistName: artistName}
	if duration != nil {
		newTrack.Duration = duration
	}
	repository.tracks = append(repository.tracks, newTrack)
	return newTrack, nil
}

func (repository *trackRepository) Delete(name string) error {
	for i, track := range repository.tracks {
		if name == track.Name {
			repository.tracks = append(repository.tracks[:i], repository.tracks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("track with name name='%s' not found", name)
}

func (repository *trackRepository) FindAll() ([]Track, error) {
	return repository.tracks, nil
}

func (repository *trackRepository) FindByName(name string) (Track, error) {
	for _, track := range repository.tracks {
		if track.Name == name {
			return track, nil
		}
	}
	return Track{}, fmt.Errorf("tracks with name='%v' not found", name)
}

func (repository *trackRepository) FindByDiscName(discName string) ([]Track, error) {
	tracks := make([]Track, 0)
	for _, track := range repository.tracks {
		if track.DiscName == discName {
			tracks = append(tracks, track)
		}
	}
	if len(tracks) == 0 {
		return tracks, fmt.Errorf("tracks with disc name='%v' not found", discName)
	}
	return tracks, nil
}
