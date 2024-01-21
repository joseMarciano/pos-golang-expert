package main

import "net/http"

func main() {
	// if we dont use our own MUX, a third part lib can attach a router in our server for example.
	// If we don't pass a mux, the http will use the global standard MUX
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", mux)
}
