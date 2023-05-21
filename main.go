package main

import (
	"SystemBaseProxy/model"
	"SystemBaseProxy/network"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// GET masterdataversion
	mux.Handle("/masterdata/version", network.MethodHandler{"GET": http.HandlerFunc(model.GetMasterDataVersion)}["GET"])

	// post playdata/updload
	mux.Handle("/playdata/upload", network.MethodHandler{"POST": http.HandlerFunc(model.PostPlayDataUpload)}["POST"])
	log.Fatal(http.ListenAndServe(":8080", mux))
}
