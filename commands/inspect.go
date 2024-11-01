package commands

import (
	"fmt"
	"pokedexcli/internal"
)

func Inspect(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error {
	if param == "" {
		fmt.Println("No param specified")
		return nil
	}

	pokemon, ok := pokedex.Caught[param]

	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s \n", pokemon.Name)
		fmt.Printf("Height: %v \n", pokemon.Height)
		fmt.Printf("Weight: %v \n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, v := range pokemon.Stats {
			fmt.Printf("    - %-15s %v \n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Stats:")
		for _, v := range pokemon.Types {
			fmt.Printf("    - %-10s \n", v.Type.Name)
		}
	}

	return nil
}
