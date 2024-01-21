package main

import (
	"encoding/json"
	"net/http"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		cep := r.URL.Query().Get("cep")
		if cep == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cepResponse, err := findCEP(cep)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(cepResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
		return
	}
}

func findCEP(cep string) (*CepResponse, error) {
	resp, err := http.Get(strings.ReplaceAll(BASE_URL, "${CEP_VALUE}", cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var cepResponse CepResponse
	err = json.NewDecoder(resp.Body).Decode(&cepResponse)
	if err != nil {
		return nil, err
	}
	return &cepResponse, nil
}
