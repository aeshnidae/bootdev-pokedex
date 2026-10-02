package pokeapi

import (
	"encoding/json"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	dat, err := c.get(url)
	if err != nil {
		return Pokemon{}, err
	}
	pokemonResp := Pokemon{}
	err = json.Unmarshal(*dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemonResp, nil
}
