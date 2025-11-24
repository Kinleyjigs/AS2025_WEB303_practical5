FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum* ./
COPY . .
RUN go mod tidy
RUN go mod download

RUN go build -o menu-service .

EXPOSE 8082

CMD ["./menu-service"]
