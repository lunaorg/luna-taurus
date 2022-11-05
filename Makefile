build:
	export GO111MODULE=on
	export CGO_ENABLED=1
	env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/link-consumer link-consumer/main.go
	env GOOS=linux go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o bin/save-images save-images/main.go
clean:
	rm -rf bin