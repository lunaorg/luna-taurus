build:
	export GO111MODULE=on
	export CGO_ENABLED=1
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/link-consumer link-consumer/main.go
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/save-images save-images/main.go

clean:
	rm -rf bin