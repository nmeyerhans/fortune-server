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

package main

import (
	"github.com/nmeyerhans/fortune-server/server"
	"log"
	"net/http"
)


func main() {
	requestCounter := make(chan int)
	errorCounter   := make(chan int)
	byteCounter    := make(chan uint64)

	go server.StatsTracker(requestCounter, errorCounter, byteCounter)
	http.Handle("/status", http.HandlerFunc(server.MakeHealthcheckFunc(requestCounter, errorCounter, byteCounter)))
	http.Handle("/", http.HandlerFunc(server.MakeServerFunc(requestCounter, errorCounter, byteCounter)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

