dev:
	air -c air.toml

build:
	go build -o tmp/main cmd/main.go
	./tmp/main


