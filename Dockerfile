# build the application in a container
FROM golang:1.14-alpine as builder
RUN mkdir /build
WORKDIR /build

# Cache layer for dependencies
ADD go.mod go.sum ./
RUN go mod download

# Cached layer for source code
ADD . .
RUN ls -a
RUN go build -o bin/anime-skip-api cmd/anime-skip-api/main.go

# Make the final image with just the built binary, excluding anything required to do the build
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/bin/anime-skip-api /app/
WORKDIR /app
EXPOSE 8081
CMD ["./anime-skip-api"]
