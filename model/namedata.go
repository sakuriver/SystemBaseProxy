package model

import (
	"SystemBaseProxy/iowriter"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type NameUpdateResponse struct {
	Result bool `json:"result"`
}

var nameUpdateRes = NameUpdateResponse{
	Result: true,
}

// 名称更新メソッド
func PutNameUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "put" {
		log.Println(errors.New("非対応メソッドです"))

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(ErrorRes{
			Message: "非対応メソッドです",
		}); err != nil {
			log.Println(err)
		}
		iowriter.ResposeWriteFromData(w, buf.String())
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(nameUpdateRes); err != nil {
		log.Fatal(err)
	}
	iowriter.ResposeWriteFromData(w, buf.String())
}
