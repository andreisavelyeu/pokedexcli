package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal"
)

func decodeAndPrintExplore(data []byte) error {
	response := internal.LocationAreaResponse{}
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&response)

	if err != nil {
		return fmt.Errorf("decoding response failed: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, v := range response.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}
	return nil
}

func Explore(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error {

	if param == "" {
		fmt.Println("No param specified")
		return nil
	}

	url := conf.BaseUrl + "/" + param

	result, ok := conf.Cache.Get(url)

	if ok {
		fmt.Println("Using cache in explore")
		return decodeAndPrintExplore(result)
	}

	fmt.Println("Making request explore")
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("response failed: %w", err)
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)

	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("reading body failed: %w", err)
	}

	conf.Cache.Add(url, bodyBytes)

	return decodeAndPrintExplore(bodyBytes)
}
