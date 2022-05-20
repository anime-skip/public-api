# build the application in a container
FROM golang:1.18-alpine as builder
# Define user for scratch image
RUN apk add --update gcc libc-dev
ENV USER=appuser
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"
# Cache layer for dependencies
WORKDIR /build
ADD go.mod go.sum ./
RUN go mod download
RUN go mod verify
# Build the binary
ADD . .
ARG VERSION
ARG STAGE
RUN : "${VERSION:?Build argument needs to be passed and non-empty}"
RUN : "${STAGE:?Build argument needs to be passed as 'development' or 'production'}"
RUN : "${VERSION:?Argument needs to be passed and non-empty.}"
RUN OOS=linux GOARCH=amd64 go build \
    -trimpath \
    -ldflags "-X main.VERSION=$VERSION -X main.STAGE=$STAGE" \
    -o bin/server \
    cmd/server/main.go

# Final image from scratch
FROM alpine
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /build/bin/server /
USER appuser:appuser
ENTRYPOINT ["/server"]
