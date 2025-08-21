# Stage 1: build
FROM golang:1.24 AS builder
WORKDIR /app

# copy go.mod & go.sum
COPY go.mod go.sum ./
RUN go mod download

# copy all source code
COPY . .

# build binary
RUN go build -o service-a main.go

# runtime
FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=builder /app/service-a .

# expose REST API port (default 8081)
EXPOSE 8081

CMD ["./service-a"]
