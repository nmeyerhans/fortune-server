package main

import (
	"github.com/nmeyerhans/hello-http/fortune"
	"encoding/json"
	"log"
	"net/http"
)

type Fortune struct {
	Message string
}

func serveFortune(w http.ResponseWriter, req *http.Request) {
	f, err := fortune.Fortune(true)
	if err != nil {
		log.Fatal(err)
	}
	m := Fortune{f}
	b, err := json.Marshal(m)
	if(err != nil) {
		log.Fatal(err)
	}
	// fmt.Print(string(b))
	w.Write(b)
}

func main() {
	http.Handle("/", http.HandlerFunc(serveFortune))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

