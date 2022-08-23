package controller

import (
	"net/http"
	"encoding/json"
	"2k22sirka/model"
	// "2k22sirka/views"
	// "log"
	// "github.com/davecgh/go-spew/spew"
	// "fmt"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			uid := r.URL.Query().Get("uid")
			data, err := model.ReadSelected(uid)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}