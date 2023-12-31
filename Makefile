dev:
	air -c air.toml

prod:
	go install github.com/a-h/templ/cmd/templ@latest
	templ generate
	go build -o tmp/main cmd/main.go
	./tmp/main


