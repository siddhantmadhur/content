package tvdb

import "github.com/siddhantmadhur/content/client"

func (t TVDB) Authenticate() bool {
	var response struct {
		Success bool `json:"success"`
	}
	err := t.Fetch(client.FetchParams{
		Endpoint: "/authentication",
	}, &response)

	if err != nil {
		return false
	}

	return response.Success
}
