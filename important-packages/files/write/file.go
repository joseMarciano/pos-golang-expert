package main

import (
	"os"
	"path/filepath"
)

func main() {
	executable, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//filepath.
	file, err := os.Create(filepath.Join(executable, "important-packages", "files", "write", "writing.txt"))
	if err != nil {
		panic(err)
	}

	file.Write([]byte("Hello World!!"))

	err = file.Close()
	if err != nil {
		return
	}
}
