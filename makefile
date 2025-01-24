build:
	go build -o app main.go

build-linux:
	SET CGO_ENABLE=0
	SET GOOS=linux
	SET GOARCH=amd64
	@echo "CGO_ENABLE=" $(CGO_ENABLE) "GOOS=" $(GOOS) "GOARCH=" $(GOARCH)
	go build -o app main.go


.PHONY: build build-linux