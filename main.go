package main

import (
	"fmt"
	"net/http"
	// "encoding/json"
	"2k22sirka/controller"
	"2k22sirka/model"
	// "log"
	// "os"
)

func main() {
	mux := controller.RegisterApi()
	db := model.ConnectDB()
	defer db.Close()
	
	fmt.Println("Serving...");
	// for heroku
	// port := os.Getenv("PORT")
	// log.Fatal(http.ListenAndServe(":"+port, mux))

	// for Local
	http.ListenAndServe(":80", mux)
}