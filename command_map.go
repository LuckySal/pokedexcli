package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	count    int     `json:"count"`
	next     *string `json:"next"`
	previous *string `"json:"previous"`
	results  []Area  `json:"results"`
}

type Area struct {
	name string `json:"name"`
	url  string `json:"url"`
}

func commandMap(config *Config) error {
	target := config.next
	if target == "" {
		target = "http://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}
	res, err := http.Get(target)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Printf("Status code: %v\n", res.StatusCode)
	if res.StatusCode > 299 {
		return fmt.Errorf("Request failed with status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", res.Body)
	area := LocationArea{}
	if err := json.Unmarshal(data, &area); err != nil {
		return err
	}
	fmt.Printf("%d\n", area.count)
	return nil
}
