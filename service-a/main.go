package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("ID  ", id)
	http.HandleFunc("/try", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Service A] HIT /try")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"dtm_result":"SUCCESS"}`))
	})

	http.HandleFunc("/cancel", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Service A] HIT /cancel (compensate)")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"dtm_result":"SUCCESS"}`))
	})

	fmt.Println("[Service A] running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
