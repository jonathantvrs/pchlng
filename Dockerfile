FROM golang:1.25.0-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /pismo ./cmd/server

FROM scratch
COPY --from=builder /pismo /pismo

EXPOSE 8080
ENTRYPOINT ["/pismo"]