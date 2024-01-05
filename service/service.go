package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/novychok/fondyapi/types"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var respose types.APIResponse
		if err := json.NewDecoder(r.Body).Decode(&respose); err != nil {
			http.Error(w, "error to encode the r.Body", http.StatusBadRequest)
			fmt.Printf("error to encode the r.Body: %+v\n", err)
			return
		}

		fmt.Println(respose)
	})

	log.Println("starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
