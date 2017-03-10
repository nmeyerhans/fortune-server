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

func logRequest(req *http.Request, responseCode int, responseSize int) {
	log.Printf("%s %s%s %d %d\n", req.RemoteAddr, req.Host, req.RequestURI, responseCode, responseSize)
}

func serveHealthcheck(w http.ResponseWriter, req *http.Request) {
	responseSize := 0
	responseCode := http.StatusOK
	defer func() {
		logRequest(req, responseCode, responseSize)
	}()
	if !fortune.Available() {
		responseCode = http.StatusInternalServerError
	}
}

func serveFortune(w http.ResponseWriter, req *http.Request) {
	responseSize := 0
	responseCode := http.StatusOK
	var body []byte

	defer func() {
		logRequest(req, responseCode, responseSize)
	}()
	fortune_text, err := fortune.Fortune(false)
	if err != nil {
		responseCode = http.StatusInternalServerError
		log.Fatal(err)
	}
	if req.Header.Get("Accept") == "application/javascript" {
		m := Fortune{fortune_text}
		body, err = json.Marshal(m)
		if(err != nil) {
			responseCode = http.StatusInternalServerError
			log.Fatal(err)
		}
	} else {
		body = []byte(fortune_text)
	}
	if req.Header.Get("UserAgent") != "" {
		fmt.Printf("Got a request from a %s browser\n", req.Header.Get("UserAgent"))
	}
	// fmt.Print(string(body))
	responseSize = len(body)
	w.Write([]byte(body))
}

func main() {
	http.Handle("/status", http.HandlerFunc(serveHealthcheck))
	http.Handle("/", http.HandlerFunc(serveFortune))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

