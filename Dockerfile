# Estágio de Build
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./
# Compila estaticamente para Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o math-api main.go

# Estágio Final (Minimalista e seguro)
FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/math-api .

EXPOSE 8000

USER 1000:1000

CMD ["./math-api"]