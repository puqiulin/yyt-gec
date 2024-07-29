.PHONY: build
build:
	env GOOS=windows GOARCH=amd64 go build -o yyt.exe .
