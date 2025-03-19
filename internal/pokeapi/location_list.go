package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ichts/pokedexcli/internal/pokecache"
)

func (c *Client) GetLocationList(pageURL *string, cache *pokecache.Cache) (locationAreaRespBody, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}
	var data locationAreaRespBody

	val, exists := cache.Get(url)
	if exists {
		err := json.Unmarshal(val, &data)
		for _, d := range data.Results {
			fmt.Println(d.Name)
		}
		if err != nil {
			return locationAreaRespBody{}, err
		}
		return data, err
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Sending request error: %v", err)
		return locationAreaRespBody{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	cache.Add(url, body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		return locationAreaRespBody{}, err
	}

	for _, d := range data.Results {
		fmt.Println(d.Name)
	}

	return data, err
}
