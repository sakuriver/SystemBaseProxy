package model

import (
	"SystemBaseProxy/network"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//
type MasterDataVersion struct {
	ID             int                 `json:"id"`
	Version        string              `json:"title"`
	RequestData    map[string][]string `json:"request_data"`
	CreatedDate    time.Time           `json:"created_date"`
	UpdateDateDate time.Time           `json:"updated_date"`
}

var version = MasterDataVersion{
	ID:             1,
	Version:        "0.01",
	CreatedDate:    time.Now(),
	UpdateDateDate: time.Now(),
}

// マスターデータのバージョン取得
func GetMasterDataVersion(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	version.RequestData = network.GetQueryParameters(r)
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&version); err != nil {
		log.Fatal(err)
	}
	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}