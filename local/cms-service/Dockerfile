FROM golang:alpine AS builder

# Set working directory
WORKDIR /app/
COPY apps/cms/ .

# Build the Go application
RUN go mod tidy
RUN /bin/sh -c 'rm -rf bin/cms-service && go build -o bin/cms-service cmd/server/main.go'
# Final stage - minimal image
FROM alpine:latest AS production

WORKDIR /app

ENV APP_ENV=production

COPY --from=builder /app/bin/cms-service ./cms-service

EXPOSE 8084

CMD ["./cms-service"]
