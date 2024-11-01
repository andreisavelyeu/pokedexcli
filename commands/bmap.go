package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal"
)

func decodeAndPrintBmap(data []byte, conf *internal.ApiConfig) error {
	response := internal.LocationAreasResponse{}
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&response)

	if err != nil {
		return fmt.Errorf("decoding response failed: %w", err)
	}

	conf.Previous = response.Previous
	conf.Next = response.Next

	for _, v := range response.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func Bmap(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error {
	url := conf.Previous

	if url == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	result, ok := conf.Cache.Get(*url)

	if ok {
		fmt.Println("Using cache in bmap")

		return decodeAndPrintBmap(result, conf)
	}

	fmt.Println("Making request bmap")
	req, err := http.NewRequest("GET", *url, nil)

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

	if err != nil {
		return fmt.Errorf("reading response body failed: %w", err)
	}

	return decodeAndPrintBmap(bodyBytes, conf)
}
