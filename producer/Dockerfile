FROM golang:1.24.1-alpine

RUN apk add --no-cache git && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN swag init -g cmd/api/main.go --parseDependency --parseInternal
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go


FROM golang:1.24.1-alpine

RUN apk --no-cache add ca-certificates

WORKDIR /app

RUN apk add --no-cache git && \
go install github.com/air-verse/air@latest

# Copy the source code and necessary files for development
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Expose port
EXPOSE 8080

# Start Air for live reloading
CMD ["air", "-c", ".air.toml"]