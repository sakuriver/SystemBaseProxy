package model

import (
	"SystemBaseProxy/iowriter"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type PlayDataResponse struct {
	Result bool `json:"result"`
}

var playDataSaveRes = PlayDataResponse{
	Result: true,
}

// プレイデータのアップロードメソッド
func PostPlayDataUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method != "post" {
		log.Fatal(errors.New("非対応メソッドです"))
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(playDataSaveRes); err != nil {
		log.Fatal(err)
	}
	iowriter.ResposeWriteFromData(w, buf.String())
}
