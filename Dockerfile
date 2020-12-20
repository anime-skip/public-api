# Default Args
ARG DEV=false

# build the application in a container
FROM golang:1.14-alpine as builder
RUN apk update
RUN apk add git
RUN mkdir /build
WORKDIR /build

# Cache layer for dependencies
ADD go.mod go.sum ./
RUN go mod download

# Cache layer for the version
ARG DEV
ADD .git .
ADD VERSION .

# Cached layer for source code
ADD . .
RUN \
  VERSION=$(cat VERSION) ;\
  if [ "$DEV" == "true" ]; then \
    SUFFIX="-$(TZ=UTC git --no-pager show --quiet --abbrev=12 --date='format-local:%Y%m%d%H%M%S' --format='%cd-%h')" ;\
  fi ;\
  go build \
    -ldflags "-X anime-skip.com/backend/internal/utils/constants.VERSION=$VERSION$SUFFIX" \
    -o bin/api \
    cmd/api/main.go

# Make the final image with just the built binary, excluding anything required to do the build
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/bin/api /app/
WORKDIR /app
EXPOSE 8081
CMD ["./api"]
