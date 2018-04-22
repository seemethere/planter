GO?=go
GOFILES:=$(shell find cli/ -name "*.go")
DEP="$(shell go env GOPATH)/bin/dep"

$(DEP): ## Grab golang/dep utility
	go get github.com/golang/dep/cmd/dep

all: build

build: bin/planter

bin/planter: cmd/planter/main.go $(GOFILES)
	mkdir -p $$(basename $@)
	$(GO) build -o $@ -v $<

.PHONY: clean
clean:
	$(RM) -r bin/

.PHONY: ensure
ensure: $(DEP)
	$< ensure -v
