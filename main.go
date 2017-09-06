package main

import (
	"encoding/json"
	"log"
	"net/http"
  "os"
  "fmt"
)

type Message struct {
	Message string `json:"message"`
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func SayHelloDrieHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("GET /say-hello-drie")

	message := Message{"Hello Zest Drie!"}

	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/say-hello-drie", SayHelloDrieHandler)
	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
