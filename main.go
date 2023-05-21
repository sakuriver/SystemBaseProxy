package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type MasterDataVersion struct {
	ID             int       `json:"id"`
	Version        string    `json:"title"`
	CreatedDate    time.Time `json:"created_date"`
	UpdateDateDate time.Time `json:"updated_date"`
}

var version = MasterDataVersion{
	ID:             1,
	Version:        "0.01",
	CreatedDate:    time.Now(),
	UpdateDateDate: time.Now(),
}

func main() {
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&version); err != nil {
			log.Fatal(err)
		}

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

	// GET /tasks
	http.HandleFunc("/masterdata/version", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
