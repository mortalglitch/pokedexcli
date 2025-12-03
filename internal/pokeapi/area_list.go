
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListAreaPokemon(selectedArea *string) (RespShallowAreas, error) {
	url := baseURL + "/location-area/" + *selectedArea
	 
	if data, found := c.Cache.Get(url); found{
		areaResp := RespShallowAreas{}
		err := json.Unmarshal(data, &areaResp)
		if err != nil {
			return RespShallowAreas{}, err
		}

		return areaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowAreas{}, err
	}
	defer resp.Body.Close()
	
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowAreas{}, err
	}

	c.Cache.Add(url, dat)

	areaResp := RespShallowAreas{}
	err = json.Unmarshal(dat, &areaResp)
	if err != nil {
		return RespShallowAreas{}, err
	}

	return areaResp, nil
}
