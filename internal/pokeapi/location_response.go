package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) ListLocations(pageURL *string) (ResponseLocationsArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, ok := client.pokeCache.Get(url)

	//If url isn't in the cache, make a http.GET and cache the answer as []byte
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ResponseLocationsArea{}, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return ResponseLocationsArea{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return ResponseLocationsArea{}, err
		}

		client.pokeCache.Add(url, data)
	}

	locationsAreaResponse := ResponseLocationsArea{}
	err := json.Unmarshal(data, &locationsAreaResponse)
	if err != nil {
		return ResponseLocationsArea{}, err
	}

	return locationsAreaResponse, nil
}
