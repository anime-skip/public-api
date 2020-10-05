# build the application in a container
FROM golang:1.14-alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN go build -o bin/anime-skip-backend cmd/anime-skip-backend/main.go

# Make the final image with just the built binary, excluding anything required to do the build
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/bin/anime-skip-backend /app/
COPY --from=builder /build/alpha.allowlist /app/alpha.allowlist
COPY --from=builder /build/test-server.allowlist /app/test-server.allowlist
WORKDIR /app
EXPOSE 8081
CMD ["./anime-skip-backend"]
