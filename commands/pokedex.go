package commands

import (
	"fmt"
	"pokedexcli/internal"
)

func Pokedex(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error {
	if len(pokedex.Caught) == 0 {
		fmt.Println("You do not have pokemons yet")
	} else {
		fmt.Println("Your Pokedex:")
		for _, v := range pokedex.Caught {
			fmt.Printf("    - %-10s \n", v.Name)
		}
	}
	return nil
}
