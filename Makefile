GOPATH:=$(shell go env GOPATH)

.PHONY: linux
linux:
		@echo "building..."
		mkdir -p bin/ && GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o ./bin/csv2tsv ./cmd/...
		$(if $(shell command -v upx), upx ./bin/csv2tsv)