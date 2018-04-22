GO?=go
GOFILES:=$(shell find commands -name "*.go")

all: build

build: bin/planter

bin/planter: cmd/planter/main.go $(GOFILES)
	mkdir -p $$(basename $@)
	$(GO) build -o $@ -v $<

.PHONY: clean
clean:
	$(RM) -r bin/
