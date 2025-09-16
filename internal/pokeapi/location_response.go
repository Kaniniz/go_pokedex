package pokeapi

import (
	"io"
	"encoding/json"
	"net/http"
)

func (client *Client) ListLocations(pageURL *string) (ResponseLocationsArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseLocationsArea{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return ResponseLocationsArea{}, err
	}
	defer resp.Body.Close()
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ResponseLocationsArea{}, err
	}

	locationsAreaResponse := ResponseLocationsArea{}
	err = json.Unmarshal(data, &locationsAreaResponse)
	if err != nil {
		return ResponseLocationsArea{}, err
	}

	return locationsAreaResponse, nil
}