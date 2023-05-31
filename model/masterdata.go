package model

import (
	"SystemBaseProxy/iowriter"
	"SystemBaseProxy/network"
	"bytes"
	"encoding/json"
	"errors"
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

	if r.Method != "get" {
		log.Fatal(errors.New("非対応メソッドです"))
		return
	}

	var buf bytes.Buffer
	version.RequestData = network.GetQueryParameters(r)
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&version); err != nil {
		log.Fatal(err)
	}
	iowriter.ResposeWriteFromData(w, buf.String())
}
