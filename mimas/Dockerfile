# Etapa 1: Construção
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Instale as dependências necessárias para compilar o código
RUN apk add --no-cache git

# Copie os arquivos de dependências e baixe os módulos
COPY go.mod go.sum ./
RUN go mod download

# Copie o restante dos arquivos do projeto
COPY . .

# Instale a ferramenta air para hot reload
RUN go install github.com/air-verse/air@latest

# Compile o binário da aplicação
RUN go build -o /app/main

# Etapa 2: Imagem mínima
FROM alpine:latest

WORKDIR /app

# Copie o binário da etapa de construção
COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env /app/.env

# Comando de inicialização da aplicação
CMD ["/app/main"]
