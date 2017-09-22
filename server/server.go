// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/nmeyerhans/fortune-server/fortune"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Fortune struct {
	Message string
}

func logRequest(req *http.Request, responseCode int, responseSize int) {
	remoteAddr := req.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr = req.RemoteAddr
	}
	log.Printf("%s %s%s %d %d\n", remoteAddr, req.Host, req.RequestURI, responseCode, responseSize)
}

func MakeHealthcheckFunc(reqCounter chan<- int, errCounter chan<- int, byteCounter chan<- uint64) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		responseSize := 0
		responseCode := http.StatusOK
		defer func() {
			reqCounter <- 1
			byteCounter <- uint64(responseSize)
			logRequest(req, responseCode, responseSize)
		}()
		if !fortune.Available(fortune.FortuneCmd) {
			errCounter <- 1
			responseCode = http.StatusInternalServerError
		}
		w.WriteHeader(responseCode)
	}
}

func MakeServerFunc(reqCounter chan<- int, errCounter chan<- int, byteCounter chan<- uint64) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		responseSize := 0
		responseCode := http.StatusOK
		var body []byte

		defer func() {
			reqCounter <- 1
			byteCounter <- uint64(responseSize)
			logRequest(req, responseCode, responseSize)
		}()
		fortune_text, err := fortune.Fortune(false)
		if err != nil {
			errCounter <- 1
			responseCode = http.StatusInternalServerError
			log.Fatal(err)
		}
		if req.Header.Get("Accept") == "application/javascript" {
			m := Fortune{fortune_text}
			body, err = json.Marshal(m)
			if(err != nil) {
				errCounter <-1
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
		w.WriteHeader(responseCode)
		w.Write([]byte(body))
	}
}

func dumpStats(requests int, errors int, bytes uint64) {
	log.Printf("Runtime stats: Total requests: %d, Failures: %d, Bytes sent: %d", requests, errors, bytes)
}

func StatsTracker(requestCounter <-chan int, failureCounter <-chan int, byteCounter <-chan uint64) {
	requests  := 0
	errors    := 0
	var bytesSent uint64 = 0
	t := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-t.C:
			go dumpStats(requests, errors, bytesSent)
		case <-requestCounter:
			requests += 1
		case <-failureCounter:
			errors += 1
		case b := <-byteCounter:
			bytesSent += b
		}
	}
}

