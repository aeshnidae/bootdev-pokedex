package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location
	dat, err := c.get(url)
	if err != nil {
		return Location{}, err
	}
	exploreResp := Location{}
	err = json.Unmarshal(*dat, &exploreResp)
	if err != nil {
		return Location{}, err
	}
	return exploreResp, nil
}
