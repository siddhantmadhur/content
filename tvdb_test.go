package content

import (
	"fmt"
	"os"
	"testing"
)

func TestTVDB(t *testing.T) {
	apiKey := os.Getenv("TVDB_API_KEY")
	readToken := os.Getenv("TVDB_READ_TOKEN")
	if apiKey == "" {
		fmt.Printf("[ERROR]: TVDB_API_KEY env variable not provided.\n")
		t.FailNow()
	}

	var tvdb TVDB
	tvdb.ApiKey = apiKey
	tvdb.ApiReadAccessToken = readToken

	var result struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}
	err := tvdb.Fetch("GET", "/movie/11", &result)
	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}

	if result.Title != "Star Wars" {
		fmt.Printf("[ERROR]: Title does not match Star Wars\n")
		t.FailNow()
	}

}
