package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"pokedexcli/internal"
)

func catchChance(pokemon internal.PokemonResponse) bool {
	baseCatchChance := 10.0

	chance := baseCatchChance / float64(pokemon.BaseExperience)

	randomChance := rand.Float64()

	return randomChance <= chance
}

func decodeAndPrintCatch(data []byte, pokedex *internal.Pokedex) error {
	response := internal.PokemonResponse{}
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&response)

	if err != nil {
		return fmt.Errorf("decoding response failed: %w", err)
	}

	caught := catchChance(response)

	if caught {
		fmt.Printf("%s was caught!\n", response.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		pokedex.Add(response)
	} else {
		fmt.Printf("%s escaped!\n", response.Name)
	}
	return nil
}

func Catch(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error {
	url := internal.PokemonUrl + "/" + param

	fmt.Printf("Throwing a Pokeball at ... %s\n", param)
	result, ok := conf.Cache.Get(url)

	if ok {
		fmt.Println("Using cache in Catch")
		return decodeAndPrintCatch(result, pokedex)
	}

	fmt.Println("Making request pokemon")
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("response failed: %w", err)
	}

	defer res.Body.Close()
	conf.Cache.Add(url, bodyBytes)
	return decodeAndPrintCatch(bodyBytes, pokedex)

}
