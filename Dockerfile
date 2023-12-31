FROM golang:1.21.2

WORKDIR /usr/local/app

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go mod tidy
RUN go build -o ./tmp/main ./cmd/main.go

ENTRYPOINT [ "tmp/main" ]
