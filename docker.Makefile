GO_PACKAGE_PATH=/go/src/github.com/seemethere/planter
GOLANG_IMAGE?=golang
GOLANG_VERSION?=1.10
GOLANG_RUN=docker run --rm -it -v "$(CURDIR)":$(GO_PACKAGE_PATH) -w "$(GO_PACKAGE_PATH)" "$(GOLANG_IMAGE):$(GOLANG_VERSION)"

shell:
	$(GOLANG_RUN) bash

%:
	$(MAKE) GO="$(GOLANG_RUN)" $@
