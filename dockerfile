FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
#instal psql 
# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh
#build go app
RUN go mod download
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/.env .
EXPOSE 8000
CMD ["./app"]
