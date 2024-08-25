.PHONY: build clean package

build:
	env GOSUMDB=off GOOS=linux GOARCH=amd64 go build -o bin/main cmd/main.go

clean:
	rm -rf ./bin

lint:
	golangci-lint run ./... --config ./build/golang-lint/config.yml

package: clean build
	cd bin && rm -f main.zip && zip main.zip main


test:
	go test ./... -v
