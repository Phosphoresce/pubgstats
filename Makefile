GOC=go build
GOFLAGS=-a -ldflags '-s'
CGOR=CGO_ENABLED=0

all: build

build:
	$(GOC) pubgstats.go

run:
	go run pubgstats.go

stat:
	$(CGOR) $(GOC) $(GOFLAGS) pubgstats.go

docker: build
	docker build .

fmt:
	gofmt -w .

clean:
	rm pubgstats
