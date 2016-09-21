hello-http: main.go
	go build

clean:
	go clean

docker: hello-http
	docker build -t hello-http .

all: docker
