package main

// go run find-cep.go [valid_cep]

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type CepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

const BASE_URL = "https://viacep.com.br/ws/${CEP_VALUE}/json/"

func main() {
	//os.Args[1:] -> get all args passed from go run find-cep.go [args]
	for _, cep := range os.Args[1:] {

		req, err := http.Get(strings.ReplaceAll(BASE_URL, "${CEP_VALUE}", cep))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			return
		}

		defer req.Body.Close()
		resp, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse: %v\n", err)
			return
		}

		var data CepResponse
		err = json.Unmarshal(resp, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse: %v\n", err)
			return
		}

		fmt.Println(data)
	}

}
