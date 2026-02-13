package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your pokedex is empty")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pm := range cfg.caughtPokemon {
		fmt.Println(" -", pm.Name)
	}
	return nil
}
