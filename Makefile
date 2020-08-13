# Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# 	You may obtain a copy of the License at

# http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

export GO111MODULE=on
export DOCKER_BUILDKIT=1

SOURCES:=$(shell find . -name '*.go' ! -name '*_test.go')
GOMOD := $(shell go env GOMOD)
# GOSUM := $(GOMOD:.mod=.sum)
DOCKER_IMAGE_TAG?=latest

.PHONY: clean docker all test

SUBDIRS:=fortune server
DOCKER_TAG:=latest

GO       = go
NAME     = fortune-server

fortune-server all: $(SOURCES) $(GOMOD) # $(GOSUM)
	$(GO) build -o $(NAME) main.go

test:
	for dir in $(SUBDIRS); do \
		( cd $$dir && go test -cover ) ; \
	done

clean:
	rm -f *~ fortune/*~ ecs/*~
	go clean

docker: fortune-server
	docker build --progress=plain --pull=true -t fortune-server:$(DOCKER_TAG) .
