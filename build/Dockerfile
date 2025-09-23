# Dockerfile (backend)
FROM golang:1.21-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /bin/app ./cmd/server

FROM gcr.io/distroless/static:nonroot
COPY --from=build /bin/app /bin/app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/bin/app"]
