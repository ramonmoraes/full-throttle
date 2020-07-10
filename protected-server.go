package main

import (
	"fmt"
	"log"
	"net/http"
)

func protectedServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", logHandler)

	fmt.Println("Protected up at", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func logHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("<-- Protected server received request"))
}
