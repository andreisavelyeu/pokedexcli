package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal"
)

func decodeAndPrint(data []byte, conf *internal.ApiConfig, url *string) error {
	response := internal.LocationAreasResponse{}
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&response)

	if err != nil {
		return fmt.Errorf("decoding response failed: %w", err)
	}

	previous := url
	if response.Previous != nil {
		previous = response.Previous
	}
	conf.Previous = previous
	conf.Next = response.Next

	for _, v := range response.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func Map(conf *internal.ApiConfig, param string, pokedex *internal.Pokedex) error {
	url := conf.Next

	if url == nil {
		base := conf.BaseUrl + "?offset=0&limit=20"
		url = &base
	}

	result, ok := conf.Cache.Get(*url)

	if ok {
		fmt.Println("Using cache in map")
		return decodeAndPrint(result, conf, url)
	}

	fmt.Println("Making request map")
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
	conf.Cache.Add(*url, bodyBytes)

	return decodeAndPrint(bodyBytes, conf, url)
}
