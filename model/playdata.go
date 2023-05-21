package model

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(playDataSaveRes); err != nil {
		log.Fatal(err)
	}
	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}
