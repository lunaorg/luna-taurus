build:
	export GO111MODULE=on
	export CGO_ENABLED=1
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/handler main.go

clean:
	rm -rf bin