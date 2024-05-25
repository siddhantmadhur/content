package tvdb

import (
	"fmt"
	"strings"

	"github.com/siddhantmadhur/content/client"
)

func (t *TVDB) SearchMovie(param client.SearchParam) (client.MovieSearchResponse, error) {

	var response client.MovieSearchResponse
	err := t.Fetch(client.FetchParams{
		Endpoint: "/search/movie",
		Queries:  []string{"query=" + strings.ReplaceAll(param.Query, " ", "%20"), fmt.Sprintf("first_air_date_year=%d", param.Year)},
	}, &response)
	if err != nil {
		return client.MovieSearchResponse{}, err
	}

	return response, nil
}

func (t *TVDB) SearchShows(param client.SearchParam) (client.ShowSearchResponse, error) {

	var response client.ShowSearchResponse
	err := t.Fetch(client.FetchParams{
		Endpoint: "/search/tv",
		Queries:  []string{"query=" + strings.ReplaceAll(param.Query, " ", "%20"), fmt.Sprintf("first_air_date_year=%d", param.Year)},
	}, &response)

	if err != nil {
		return client.ShowSearchResponse{}, err
	}

	return response, nil
}
