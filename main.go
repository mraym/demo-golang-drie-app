package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func SayHelloDrieHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("GET /say-hello-drie")

	message := Message{"Hello Golang Drie!"}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {

	http.HandleFunc("/say-hello-drie", SayHelloDrieHandler)
	http.ListenAndServe(":8088", nil)
}
