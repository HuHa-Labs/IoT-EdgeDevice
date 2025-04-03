FROM golang:1.24.2-alpine3.21
LABEL authors="phuongdoantuminh"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "."]