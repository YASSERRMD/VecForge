FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY cmd/vecforge-server go.* ./
RUN go mod download
COPY . .
RUN go build -o vecforge .

FROM alpine
WORKDIR /app
COPY --from=builder /app/vecforge .
EXPOSE 8080
CMD ["./vecforge"]
