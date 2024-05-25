package tvdb

import (
	"fmt"
	"os"
	"testing"

	"github.com/siddhantmadhur/content/types"
)

func TestSearchMovie(t *testing.T) {
	readToken := os.Getenv("TVDB_READ_TOKEN")
	if readToken == "" {
		fmt.Printf("[ERROR]: TVDB_READ_TOKEN env variable not provided.\n")
		t.FailNow()
	}

	var tvdb TVDB
	tvdb.ApiKey = readToken

	res, err := tvdb.SearchMovies(types.SearchParam{
		Query: "star wars empire strikes back",
	})
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}

	if len(res.Results) == 0 {
		fmt.Printf("[ERROR]: Did not get any results\n")
		t.FailNow()
	}

	if res.Results[0].Title != "The Empire Strikes Back" {
		fmt.Printf("[ERROR]: Name does not match\n")
		t.FailNow()
	}
}

func TestSearchShows(t *testing.T) {
	readToken := os.Getenv("TVDB_READ_TOKEN")
	if readToken == "" {
		fmt.Printf("[ERROR]: TVDB_READ_TOKEN env variable not provided.\n")
		t.FailNow()
	}

	var tvdb TVDB
	tvdb.ApiKey = readToken

	res, err := tvdb.SearchShows(types.SearchParam{
		Query: "modern family",
	})
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}

	if len(res.Results) == 0 {
		fmt.Printf("[ERROR]: Did not get any results\n")
		t.FailNow()
	}

}
