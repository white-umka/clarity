FROM golang:1.26.7-alpine
WORKDIR /app
COPY . .
RUN go build -o bot ./cmd/bot
CMD ["./bot"]
