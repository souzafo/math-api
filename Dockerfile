# Estágio de Build (Adaptável à plataforma de destino)
FROM --platform=$BUILDPLATFORM golang:alpine AS builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./
# Compila estaticamente para o SO e a Arquitetura de destino
RUN CGO_ENABLED=0 GOOS=\({TARGETOS:-linux} GOARCH=\){TARGETARCH:-amd64} go build -o math-api main.go

# Estágio Final (Imagem minimalista e segura)
FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/math-api .

EXPOSE 8000

USER 1000:1000

CMD ["./math-api"]