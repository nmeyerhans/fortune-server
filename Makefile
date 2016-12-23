hello-http: main.go
	go build

clean:
	rm -f *~ fortune/*~ ecs/*~
	go clean

docker: hello-http
	docker build -t hello-http .

all: docker
