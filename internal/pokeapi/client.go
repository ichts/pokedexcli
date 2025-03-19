package pokeapi

import (
	"net/http"
)

type Client struct {
	HttpClient http.Client
}
