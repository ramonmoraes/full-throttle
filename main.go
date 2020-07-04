package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", logHandler)

	fmt.Println("Listening")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func logHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received")
	w.Write([]byte("reken\n"))
	w.Header().Set("Content-Type", "text/plain")
}
