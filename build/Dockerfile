# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /forya-be ./cmd/main

# Final image
FROM gcr.io/distroless/static:nonroot
COPY --from=builder /forya-be /forya-be
USER nonroot
EXPOSE 8080
ENTRYPOINT ["/forya-be"]
