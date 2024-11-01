package main

import (
	"pokedexcli/commands"
	"pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error
}

var commandsMap = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commands.Help,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commands.Exit,
	},
	"map": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commands.Map,
	},
	"bmap": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commands.Bmap,
	},
	"explore": {
		name:        "explore",
		description: "Explores the Pokemon",
		callback:    commands.Explore,
	},
	"catch": {
		name:        "catch",
		description: "Catches the Pokemon",
		callback:    commands.Catch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspects your pokemon",
		callback:    commands.Inspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Lists all your pokemons",
		callback:    commands.Pokedex,
	},
}

func getCommand(command string) (cliCommand, bool) {
	res, ok := commandsMap[command]
	return res, ok
}
