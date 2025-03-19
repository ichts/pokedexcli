package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ichts/pokedexcli/internal/pokecache"
)

func (c *Client) GetPokemonList(location string, cache *pokecache.Cache) ([]string, error) {
	url := baseLocationURL + "/" + location

	var data exploreRespBody
	var pokemonList []string

	val, exists := cache.Get(location)
	if exists {
		err := json.Unmarshal(val, &pokemonList)
		return pokemonList, err
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Sending request error: %v", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("decoding json error: %v\n", err)
		return nil, err
	}
	for _, p := range data.PokemonEncounters {
		pokemonList = append(pokemonList, p.Pokemon.Name)
	}

	pokemonListJSON, err := json.Marshal(pokemonList)
	if err != nil {
		fmt.Printf("encoding json error: %v\n", err)
		return nil, err

	}
	cache.Add(location, pokemonListJSON)

	return pokemonList, err
}
