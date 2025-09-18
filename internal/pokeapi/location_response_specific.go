package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) ListLocationSpecific(areaName string) (ResponseLocationAreaSpecific, error) {
	url := baseURL + "/location-area/" + areaName
	data, ok := client.pokeCache.Get(url)

	//If url isn't in the cache, make a http.GET and cache the answer as []byte
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ResponseLocationAreaSpecific{}, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return ResponseLocationAreaSpecific{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return ResponseLocationAreaSpecific{}, err
		}

		client.pokeCache.Add(url, data)
	}

	locationResponseSpecific := ResponseLocationAreaSpecific{}
	err := json.Unmarshal(data, &locationResponseSpecific)
	if err != nil {
		return ResponseLocationAreaSpecific{}, err
	}

	return locationResponseSpecific, nil
}
