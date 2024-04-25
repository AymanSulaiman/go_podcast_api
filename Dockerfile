FROM golang:1.22.2

WORKDIR /app

# Copy go mod and sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

CMD ["go","run","main.go"]