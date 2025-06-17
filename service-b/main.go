package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("ID  ", id)
	http.HandleFunc("/try", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Service B] HIT /try")
		var payload struct {
			Amount int `json:"amount"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, `{"dtm_result":"FAILURE","message":"invalid request"}`, http.StatusBadRequest)
			log.Println("Failed to decode JSON:", err)
			return
		}

		log.Printf("Received amount: %d\n", payload.Amount)

		// Simulasi kegagalan jika amount > 500
		if payload.Amount > 500 {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"dtm_result":"FAILURE","message":"simulated failure"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"dtm_result":"SUCCESS"}`))

	})

	http.HandleFunc("/cancel", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Service B] HIT /cancel")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"dtm_result":"SUCCESS"}`))
	})

	fmt.Println("[Service B] running on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
