package content

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TVDB struct {
	ApiKey             string `json:"api_key"`
	ApiReadAccessToken string `json:"api_read_access_token"`
}

func (t *TVDB) Fetch(method string, endpoint string, result any) error {
	client := &http.Client{}

	req, err := http.NewRequest(method, fmt.Sprintf("https://api.themoviedb.org/3%s?api_key=%s", endpoint, t.ApiKey), nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(result)

	return err
}
