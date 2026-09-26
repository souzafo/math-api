# Estágio de Build
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o math-api main.go

# Estágio Final
FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/math-api .

EXPOSE 8000

USER 1000:1000

CMD ["./math-api"]