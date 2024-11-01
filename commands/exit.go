package commands

import (
	"fmt"
	"os"
	"pokedexcli/internal"
)

func Exit(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error {
	fmt.Println("Exiting")
	os.Exit(0)
	return nil
}
