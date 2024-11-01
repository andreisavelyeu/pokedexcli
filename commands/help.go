package commands

import (
	"fmt"
	"pokedexcli/internal"
)

func Help(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}
