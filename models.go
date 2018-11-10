package golang_graphql

type Artist struct {
	Name  string `json:"name"`
	Genre *Genre `json:"genre"`
}
