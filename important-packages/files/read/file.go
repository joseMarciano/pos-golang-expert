package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	executable, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//filepath.
	file, err := os.ReadFile(filepath.Join(executable, "important-packages", "files", "read", "example.txt"))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	if err != nil {
		return
	}
}
