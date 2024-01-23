package main

import (
	"io"
	"net/http"
)

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {

		panic(err)
	}
	defer resp.Body.Close()

	value, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	print(string(value))
}
