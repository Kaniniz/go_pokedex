package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(pokemon string) (PokemonInformation, error) {
	url := baseURL + "/pokemon/" + pokemon

	data, ok := client.pokeCache.Get(url)

	//If url isn't in the cache, make a http.GET and cache the answer as []byte
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonInformation{}, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return PokemonInformation{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PokemonInformation{}, err
		}

		client.pokeCache.Add(url, data)
	}

	pokemonInformation := PokemonInformation{}
	err := json.Unmarshal(data, &pokemonInformation)
	if err != nil {
		return PokemonInformation{}, err
	}

	return pokemonInformation, nil
}