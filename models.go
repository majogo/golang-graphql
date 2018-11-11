package golang_graphql

type Artist struct {
	Name  string `json:"name"`
	Genre *Genre `json:"genre"`
}

type Disc struct {
	ArtistName       string  `json:"artistName"`
	Name             string  `json:"name"`
	Year             int     `json:"year"`
}

type Track struct {
	ArtistName string `json:"artistName"`
	DiscName   string `json:"discName"`
	Name       string `json:"name"`
	Duration   *int   `json:"duration"`
}
