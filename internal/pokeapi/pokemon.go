package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokemon(name string) Pokemon {
	url := basePokemonURL + "/" + name

	resp, err := c.HttpClient.Get(url)
	if err != nil {
		fmt.Printf("Sending request error: %v", err)
	}
	defer resp.Body.Close()

	var pm Pokemon

	err = json.NewDecoder(resp.Body).Decode(&pm)
	if err != nil {
		fmt.Printf("decoding json error: %v", err)
	}

	return pm
}
