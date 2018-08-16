package main

import (
	"net/http"
	"encoding/json"
	"io"
	"../core"
)

var blockchain *core.Blockchain

func run()  {
	http.HandleFunc("/blockchain/get", blockchianGetHandler)
	http.HandleFunc("/blockchain/write", blockchianWriteHandler)
	http.ListenAndServe("localhost:4010", nil)
}

func blockchianGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchianWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchianGetHandler(w, r)
}

func main () {
	blockchain = core.NewBlockchain()
	run()
}