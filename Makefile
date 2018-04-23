GO?=go
GOARCH=$(shell go env GOARCH)
GOOS=$(shell go env GOOS)
GOFILES:=$(shell find cli/ -name "*.go")
DEP="$(shell go env GOPATH)/bin/dep"

all: build

$(DEP): ## Grab golang/dep utility
	go get github.com/golang/dep/cmd/dep

build: bin/$(GOARCH)/$(GOOS)/planter

bin/$(GOARCH)/$(GOOS)/planter: cmd/planter/main.go $(GOFILES)
	@mkdir -p $$(basename $@)
	$(GO) build -o $@ -v $<

.PHONY: clean
clean:
	$(RM) -r bin/

.PHONY: ensure
ensure: $(DEP)
	$< ensure -v
