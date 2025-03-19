# Stage 1: Build
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . . 

RUN go build -o main .

# Final stage
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/main .

CMD ["./main"]