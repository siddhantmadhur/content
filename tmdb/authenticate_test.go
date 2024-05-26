package tmdb

import (
	"fmt"
	"os"
	"testing"
)

func TestAuthenticate(t *testing.T) {

	readToken := os.Getenv("TVDB_READ_TOKEN")
	if readToken == "" {
		fmt.Printf("[ERROR]: TVDB_READ_TOKEN env variable not provided.\n")
		t.FailNow()
	}

	var tvdb Client
	tvdb.ApiKey = readToken

	isAuthenticated := tvdb.Authenticate()
	if !isAuthenticated {
		fmt.Printf("[ERROR] Could not authenticate user\n")
		t.FailNow()
	}

}
