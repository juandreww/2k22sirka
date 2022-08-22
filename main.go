package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

		}
	})
	// defer db.Close()
	
	fmt.Println("Serving...");
	http.ListenAndServe(":80", mux)
}