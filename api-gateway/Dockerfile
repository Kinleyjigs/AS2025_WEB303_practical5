FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum* ./
COPY . .
RUN go mod tidy
RUN go mod download

RUN go build -o api-gateway .

EXPOSE 8080

CMD ["./api-gateway"]
