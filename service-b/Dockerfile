# Stage 1: build
FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o service-b main.go

# Stage 2: runtime
FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=builder /app/service-b .

EXPOSE 8080
EXPOSE 50051

CMD ["./service-b"]
