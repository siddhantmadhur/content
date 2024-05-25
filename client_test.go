package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/siddhantmadhur/content/tvdb"
)

func TestTVDBClient(t *testing.T) {
	client, err := NewClient(tvdb.TVDB{
		ApiKey: os.Getenv("TVDB_READ_TOKEN"),
	})

	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err.Error())
		t.FailNow()
	}

	authenticate := client.Authenticate()

	if !authenticate {
		fmt.Printf("[ERROR]: Could not authenticate client\n")
		t.FailNow()
	}
}
