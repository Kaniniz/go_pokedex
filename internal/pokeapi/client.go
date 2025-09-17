package pokeapi

import (
	"net/http"
	"time"

	"github.com/Kaniniz/go_Pokedex/internal/pokecache"
)

type Client struct {
	pokeCache  pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheTime time.Duration) Client {
	return Client{
		pokeCache: pokecache.NewCache(cacheTime),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}

}
