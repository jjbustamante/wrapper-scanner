BINARY_NAME=scanner
OUT_DIR?=$(PWD)/out

all: build

build:
	mkdir -p $(OUT_DIR)
	GOARCH=amd64 GOOS=darwin go build -o $(OUT_DIR)/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o $(OUT_DIR)/${BINARY_NAME}-linux main.go

clean:
	go clean
	rm -R $(OUT_DIR)
