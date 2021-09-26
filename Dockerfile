# build the application in a container
FROM golang:1.14-alpine as builder
WORKDIR /build

# Cache layer for dependencies
ADD go.mod go.sum ./
RUN go mod download

# Cached layer for source code
ADD . .
ARG VERSION
RUN : "${VERSION:?Build argument needs to be passed and non-empty.}"
RUN \
  go build \
    -ldflags "-X anime-skip.com/backend/internal/utils/constants.VERSION=$VERSION" \
    -o bin/api-service \
    cmd/api-service/main.go

# Make the final image with just the built binary, excluding anything required to do the build
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
WORKDIR /app
COPY --from=builder /build/bin/api-service .
CMD ["./api-service"]
