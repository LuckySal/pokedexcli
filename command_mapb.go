package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(config *Config) error {
	target := config.previous
	if target == "" {
		fmt.Printf("You're on the first page\n")
		return nil
	}
	res, err := http.Get(target)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Request failed with status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	area := LocationArea{}
	json.Unmarshal(data, &area)
	fmt.Printf("%d\n", area.count)
	return nil
}
