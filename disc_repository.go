package golang_graphql

import "fmt"

type discRepository struct {
	discs []Disc
}

var discRepositoryInstance discRepository

func init() {
	discRepositoryInstance = discRepository{
		discs: make([]Disc, 0),
	}
	_, err := discRepositoryInstance.Add("U2", "The Joshua Tree", 1987)
	if err != nil {
		fmt.Println(fmt.Printf("Error during adding U2 - The Joshua Tree to discs: '%v'", err))
	}
	_, err = discRepositoryInstance.Add("U2", "The Unforgettable Fire", 1984)
	if err != nil {
		fmt.Println(fmt.Printf("Error during adding U2- The Unforgettable Fire to discs: '%v'", err))
	}
	_, err = discRepositoryInstance.Add("Snap", "The Madman's Return", 1992)
	if err != nil {
		fmt.Println(fmt.Printf("Error during adding Snap- The Madman's Return: '%v'", err))
	}
}

func (repository *discRepository) Add(artistName, name string, year int) (Disc, error) {
	newDisc := Disc{ArtistName: artistName, Name: name, Year: year}
	repository.discs = append(repository.discs, newDisc)
	return newDisc, nil
}

func (repository *discRepository) Delete(name string) error {
	for i, disc := range repository.discs {
		if name == disc.Name {
			repository.discs = append(repository.discs[:i], repository.discs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("disc with name name='%s' not found", name)
}

func (repository *discRepository) FindAll() ([]Disc, error) {
	return repository.discs, nil
}

func (repository *discRepository) FindByName(name string) (Disc, error) {
	for _, disc := range repository.discs {
		if disc.Name == name {
			return disc, nil
		}
	}
	return Disc{}, fmt.Errorf("disc with name='%v' not found", name)
}
func (repository *discRepository) FindByArtistName(artistName string) ([]Disc, error) {
	discs := make([]Disc, 0)
	for _, disc := range repository.discs {
		if disc.ArtistName == artistName {
			discs = append(discs, disc)
		}
	}
	return discs, nil
}
