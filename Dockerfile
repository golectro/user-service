FROM golang:1.24.4-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/web

EXPOSE 8080

CMD ["./app"]
