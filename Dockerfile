# ---- Build stage ----
FROM golang:1.24 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ---- Runtime stage ----
FROM alpine:3.20

WORKDIR /app

# Copy compiled binary
COPY --from=builder /app/server .

# Copy only static assets (without photos)
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./server"]
