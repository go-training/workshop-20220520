GOFMT ?= gofumpt -l -w -extra
GO ?= go
GOFILES := $(shell find . -name "*.go")

.PHONY: fmt
fmt:
	@which gofumpt > /dev/null; if [ $$? -ne 0 ]; then \
		$(GO) install mvdan.cc/gofumpt@latest; \
	fi
	$(GOFMT) -w $(GOFILES)
