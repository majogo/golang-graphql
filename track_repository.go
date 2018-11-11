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
	_, err := trackRepositoryInstance.Add("Where The Streets Have No Name", "The Joshua Tree", "U2", nil)
	if err != nil {
		fmt.Println(fmt.Printf("Error during adding U2 - Where the streets have no name to tracks: '%v'", err))
	}
	durationImStillHaventFound := 277
	_, err = trackRepositoryInstance.Add("I Still Haven't Found What I'm Looking For", "The Joshua Tree", "U2", &durationImStillHaventFound)
	if err != nil {
		fmt.Println(fmt.Printf("Error during adding U2 - I still haven't found to tracks: '%v'", err))
	}
	_, err = trackRepositoryInstance.Add("The Unforgettable Fire", "The Unforgettable Fire", "U2", nil)
	if err != nil {
		fmt.Println(fmt.Printf("Error during adding U2 - The Unforgettable Fire to tracks: '%v'", err))
	}
	durationRhythmIsADancer := 332
	_, err = trackRepositoryInstance.Add("Rhythm Is A Dancer", "The Madman's Return", "Snap", &durationRhythmIsADancer)
	if err != nil {
		fmt.Println(fmt.Printf("Error during adding Snap - Rhythm Is A Dancer to tracks: '%v'", err))
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
