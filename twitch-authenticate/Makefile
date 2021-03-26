BIN := bin
OUTPUT := twitch-authenticate

.PHONY : get build clean fmt lint test

build-linux:
	echo "Compiling for Linux"
	GOOS=linux go build -v -o ${BIN}/${OUTPUT} .

build-windows:
	echo "Compiling for Windows X64"
	GOOS=windows GOARCH=amd64 go build -v -o ${BIN}/${OUTPUT}.exe .	

build: dep build-linux build-windows

clean:
	go clean
	rm -rf ${BIN}

fmt:
	go fmt ./...

get:
	go get -v

dep:
	go mod download -x

lint:
	go get -u golang.org/x/lint/golint
	go vet ./...
	golint -set_exit_status ./...

test:
	go test -v -cover ./...