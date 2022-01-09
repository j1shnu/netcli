BIN_NAME="netcli"

build:
	echo -e "Compiling for your OS and platform\n"
	go build -o bin/${BIN_NAME} .

run:
	go run .

compile:
	echo -e "Compiling NetCLI for every OS and platform\n"
	GOOS=linux GOARCH=arm go build -o bin/${BIN_NAME}-linux-arm .
	GOOS=linux GOARCH=arm64 go build -o bin/${BIN_NAME}-linux-arm64 .
	GOOS=freebsd GOARCH=386 go build -o bin/${BIN_NAME}-freebsd-386 .
	GOOS=linux GOARCH=amd64 go build -o bin/${BIN_NAME}-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/${BIN_NAME}-darwin-amd64 .
	GOOS=windows GOARCH=amd64 go build -o bin/${BIN_NAME}-windows-amd64.exe .

clean:
	go clean
	rm bin/netcli*
