FROM golang:alpine AS builder

# Set working directory
WORKDIR /app/
COPY apps/post/ .

# Build the Go application
RUN go mod tidy
RUN /bin/sh -c 'rm -rf bin/post-service && go build -o bin/post-service cmd/server/main.go'
# Final stage - minimal image
FROM alpine:latest AS production

WORKDIR /app

ENV APP_ENV=production

COPY --from=builder /app/bin/post-service ./post-service

EXPOSE 8084

CMD ["./post-service"]
