package client

type CreateClientParams struct {
	Type   string `json:"type"`
	ApiKey string `json:"api_key"`
}

type Client interface {
	Fetch(FetchParams, any) error
	SearchMovies(SearchParam) (MovieSearchResponse, error)
	SearchShows(SearchParam) (ShowSearchResponse, error)
	Authenticate() bool
}
