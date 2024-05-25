package main

import "github.com/siddhantmadhur/content/types"

type Client interface {
	Fetch(types.FetchParams, any) error
	SearchMovies(types.SearchParam) (types.MovieSearchResponse, error)
	SearchShows(types.SearchParam) (types.ShowSearchResponse, error)
	Authenticate() bool
}
