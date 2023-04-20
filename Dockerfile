FROM golang:1.19 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o main cmd/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]