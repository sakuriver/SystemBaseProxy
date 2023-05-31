package iowriter

import (
	"fmt"
	"net/http"
)

func ResposeWriteFromData(w http.ResponseWriter, bufStr string) {
	_, err := fmt.Fprint(w, bufStr)
	if err != nil {
		return
	}
}
