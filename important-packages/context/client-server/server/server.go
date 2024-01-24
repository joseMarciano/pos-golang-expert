package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	log.Println("Request started")
	defer log.Println("Request finished")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request processed")
		writer.Write([]byte("Request processed"))
		return
	case <-ctx.Done():
		log.Println("Request cancelled by client: " + ctx.Err().Error())
		http.Error(writer, "Request cancelled by client: "+ctx.Err().Error(), http.StatusRequestTimeout)
		return
	}
}
