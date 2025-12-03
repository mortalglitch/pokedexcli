package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) GetPokemonData(selectedPokemon *string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + *selectedPokemon
	 
	if data, found := c.Cache.Get(url); found{
		pokeResp := Pokemon{}
		err := json.Unmarshal(data, &pokeResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokeResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.Cache.Add(url, dat)

	pokeResp := Pokemon{}
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokeResp, nil
}
