
package pokeapi

import (
	"net/http"
	"time"
	"github.com/mortalglitch/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	Cache      *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	cache := pokecache.NewCache(5 * time.Minute)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: cache,
	}
}
