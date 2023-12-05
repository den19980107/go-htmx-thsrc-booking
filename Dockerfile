FROM golang:latest
WORKDIR /web
COPY . .
RUN go mod download
RUN env GOOS=linux GOARCH=amd64 go build -o web ./cmd/web/main.go
EXPOSE 8000
CMD ["./web"]
