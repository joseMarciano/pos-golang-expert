package main

import (
	"bufio"
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
	file, err := os.Open(filepath.Join(executable, "important-packages", "files", "read", "read-chunks", "1MB-FILE.txt"))
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}
}
