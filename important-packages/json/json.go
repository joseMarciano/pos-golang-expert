package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BankingAccount struct {
	Name    string
	Account string
}

type Investment struct { // using tags to binding props
	Name         string  `json:"name"`
	Account      string  `json:"account"`
	Amount       float64 `json:"-"` // ignoring amount for decoding
	Currency     string  `json:"currency"`
	Type         string
	StartDate    string
	EndDate      string
	Frequency    string
	InterestRate float64
}

func main() {
	bankingAccount := BankingAccount{
		Name:    "John Doe",
		Account: "123-456-789",
	}
	bytes, err := json.Marshal(bankingAccount)
	if err != nil {
		panic(err)
	}
	println(string(bytes)) // {"Name":"John Doe","Account":"123-456-789"

	/**
	Another way to encode a json in an output
	*/
	encoder := json.NewEncoder(os.Stdout) // The value will be printed in std out
	encoder.Encode(bankingAccount)

	/**
	Another way to encode a json in a file
	*/
	file, err := os.Create("/home/usuario/JOSEPAULO/pos-golang-expert/important-packages/json/example.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder = json.NewEncoder(file) // The value will be printed in std out
	encoder.Encode(bankingAccount)

	/**
	Unmarshal
	*/

	jsonBytes := []byte(`{"Name":"John Doe","Account":"123-456-789"}`)
	var bankingAccount2 BankingAccount
	err = json.Unmarshal(jsonBytes, &bankingAccount2)
	if err != nil {
		panic(err)
	}

	fmt.Println(bankingAccount2) // John Doe

	//using tags
	investmentByts := []byte(`{"name":"John Doe","account":"123-456-789","amount":1000,"currency":"USD","type":"Fixed","startDate":"2020-01-01","endDate":"2020-01-31","frequency":"Monthly","interestRate":0.1}`)
	var investments Investment
	err = json.Unmarshal(investmentByts, &investments)
	if err != nil {
		panic(err)
	}
	fmt.Println(investments)

	/*
		Decoding from a file
	*/
	file, err = os.Open("/home/usuario/JOSEPAULO/pos-golang-expert/important-packages/json/example.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file) // The value will be printed in std out
	var bankingAccount3 BankingAccount
	decoder.Decode(&bankingAccount3)
	fmt.Println(bankingAccount2) // John Doe
	//encoder.Encode(bankingAccount)

}
