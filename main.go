package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := controller.RegisterApi()
	db := model.ConnectDB()
	defer db.Close()
	
	fmt.Println("Serving...");
	http.ListenAndServe(":80", mux)
}