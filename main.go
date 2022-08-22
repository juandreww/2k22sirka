package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"2k22sirka/structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response {
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	// defer db.Close()
	
	fmt.Println("Serving...");
	http.ListenAndServe(":80", mux)
}