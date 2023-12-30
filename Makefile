dev:
	air -c air.toml

prod:
	go build -o tmp/main cmd/main.go
	./tmp/main


