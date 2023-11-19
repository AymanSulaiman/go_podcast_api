FROM golang:1.21.4-alpine

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

EXPOSE 8000

CMD ["go","run", "main.go"]
