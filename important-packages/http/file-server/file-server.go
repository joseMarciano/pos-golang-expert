package main

import (
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// go run file-serverconfig.go
	pwd, _ := filepath.Abs("./public")
	log.Println(pwd)
	fileServer := http.FileServer(http.Dir(pwd))
	mux := http.NewServeMux()

	mux.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
