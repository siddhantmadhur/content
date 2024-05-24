package content

import "errors"

type ApiParam struct {
	Key         string `json:"key"`
	AccessToken string `json:"access_token"`
}

type NewClientParmas struct {
	TVDB ApiParam `json:"tvdb"`
}

type ApiKey struct {
	Key         string `json:"key"`
	AccessToken string `json:"access_token"`
	Url         string `json:"url"`
}

type Client struct {
	Clients []ApiKey `json:"clients"`
}

// q - Search query
func (c *Client) Search(q string) error {
	return nil
}

func NewClient(params NewClientParmas) (Client, error) {

	var client Client

	if params.TVDB.Key != "" {
		client.Clients = append(client.Clients, ApiKey{
			AccessToken: params.TVDB.AccessToken,
			Key:         params.TVDB.Key,
			Url:         "https://api.themoviedb.org/3",
		})
	}

	if len(client.Clients) == 0 {
		return client, errors.New("No provider supported")
	}

	return client, nil
}
