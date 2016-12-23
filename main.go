package main

import (
	"github.com/nmeyerhans/hello-http/fortune"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Fortune struct {
	Message string
}

func serveHealthcheck(w http.ResponseWriter, req *http.Request) {
	if !fortune.Available() {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
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
	if req.Header.Get("UserAgent") != "" {
		fmt.Printf("Got a request from a %s browser\n", req.Header.Get("UserAgent"))
	}
	// fmt.Print(string(b))
	w.Write(b)
}

func main() {
	http.Handle("/status", http.HandlerFunc(serveHealthcheck))
	http.Handle("/", http.HandlerFunc(serveFortune))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

