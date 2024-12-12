FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o library-example-api

FROM alpine
WORKDIR /folder
COPY --from=builder /app/library-example-api ./binary
ENTRYPOINT ["./binary"]
