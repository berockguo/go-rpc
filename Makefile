.PHONY: cmd build image

# Output prefix, defaults to local directory if not specified
ifeq ($(PREFIX),)
  PREFIX := $(shell pwd)
endif

BUILDTAGS ?=
GO_LDFLAGS ?=
GO_GCFLAGS ?=
VERBOSE_GO :=
GO := godep go

# Honor debug
ifeq ($(DEBUG),true)
	# Disable function inlining and variable registerization
	GO_GCFLAGS := -gcflags "-N -l"
else
	# Turn of DWARF debugging information and strip the binary otherwise
	GO_LDFLAGS := $(GO_LDFLAGS) -w -s
endif

# Honor static
ifeq ($(STATIC),true)
	# Append to the version
	GO_LDFLAGS := $(GO_LDFLAGS) -extldflags -static
endif

# Honor verbose
ifeq ($(VERBOSE),true)
	VERBOSE_GO := -v
endif

# Full package list
PKGS := $(shell go list -tags "$(BUILDTAGS)" ./... | grep -v "/Godeps/" | grep -v "/cmd")

# go files of package
PKG_GOFILE := $(shell find ./pkg -type f -name '*.go')
PKG_GOFILE_WITHOUT_TEST := $(filter-out %_test.go, $(PKG_GOFILE))

SERVICENAME := $(service)

default: build
############## build ##################
build: server client

server: build-server
build-server: $(PREFIX)/bin/$(SERVICENAME)server
$(PREFIX)/bin/$(SERVICENAME)server: $(PREFIX)/cmd/$(SERVICENAME)server.go $(PKG_GOFILE_WITHOUT_TEST)
	$(GO) build -o $@ $(VERBOSE_GO) -tags "$(BUILDTAGS)" -ldflags "$(GO_LDFLAGS)" $(GO_GCFLAGS) $<

client: build-client
build-client: $(PREFIX)/bin/$(SERVICENAME)client
$(PREFIX)/bin/$(SERVICENAME)client: $(PREFIX)/cmd/$(SERVICENAME)client.go $(PKG_GOFILE_WITHOUT_TEST)
	$(GO) build -o $@ $(VERBOSE_GO) -tags "$(BUILDTAGS)" -ldflags "$(GO_LDFLAGS)" $(GO_GCFLAGS) $<


############## test ##################
test: test-short test-long

# Quick test. You can bypass long tests using: `if testing.Short() { t.Skip("Skipping in short mode.") }`
test-short:
	$(GO) test $(VERBOSE_GO) -test.short -tags "$(BUILDTAGS)" $(PKGS)

# Runs long tests also, plus race detection
test-long:
	$(GO) test $(VERBOSE_GO) -race -tags "$(BUILDTAGS)" $(PKGS)


############## validate ##################
fmt:
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v Godeps/ | tee /dev/stderr)"

# Vet
vet: build
	@test -z "$$(go vet $(PKGS) 2>&1 | tee /dev/stderr)"


clean:
	$(RM) -r $(PREFIX)/bin
