package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal"
	"strings"
)

func main() {
	base := internal.BaseUrl
	apiConfig := &internal.ApiConfig{
		BaseUrl:  base,
		Next:     nil,
		Previous: nil,
		Cache:    internal.NewCache(),
	}

	pokedex := internal.NewPokedex()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please type a command")
	for {
		fmt.Print("> ")

		scanner.Scan()
		text := scanner.Text()

		splittedInput := strings.Split(text, " ")

		command, ok := getCommand(splittedInput[0])

		var param string

		if len(splittedInput) > 1 {
			param = splittedInput[1]
		}

		if !ok {
			fmt.Println("Command does not exist")
			continue
		}

		err := command.callback(apiConfig, param, pokedex)
		fmt.Println(err.Error())
	}
}
