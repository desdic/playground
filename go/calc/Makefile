.PHONY: all benchmark cover test build clean
default: all

GOC=go
GOBIN=$(shell pwd)

all: bin/calc

test:
	$(GOC) test -race $(FLAGS) -v internal/evaluate/*.go

benchmark:
	$(GOC) test $(FLAGS) -bench=. internal/evaluate/*.go

bin/calc: cmd/calc/*.go internal/*/*.go
	GOBIN=${GOBIN}/bin $(GOC) install -mod=vendor -buildmode=pie cmd/calc/main.go

clean:
	rm -f bin

