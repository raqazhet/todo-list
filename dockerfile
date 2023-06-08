FROM golang 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/main.go
EXPOSE 8000
CMD ["./app"]