package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/novychok/fondyapi/types"
)

func main() {
	// This callback is implemented to notify the service server about payment
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		var response types.CallbackResponse
		if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
			http.Error(w, "error to encode the r.Body", http.StatusBadRequest)
			fmt.Printf("error to encode the r.Body: %+v\n", err)
			return
		}

		fmt.Printf("%+v\n", response)
	})

	// This response is implemented to redirect the user on service webpage with final payment response
	http.HandleFunc("/response", func(w http.ResponseWriter, r *http.Request) {
		var response types.WriteResponse
		res, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		if err = json.Unmarshal([]byte(res), &response); err != nil {
			fmt.Println("unmarshal error", err)
		}

		w.Write(res)
	})

	log.Println("starting a server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
