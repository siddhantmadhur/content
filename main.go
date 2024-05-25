package main

import (
	"github.com/siddhantmadhur/content/client"
)

func NewClient(client client.Client) (client.Client, error) {
	return client, nil
}
