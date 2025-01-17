# Etapa 1: Build
FROM golang:1.23 AS builder

# Configurar o diretório de trabalho no container
WORKDIR /app

# Copiar os arquivos do projeto para o container
COPY go.mod ./
RUN go mod download

# Copiar o restante do código
COPY . .

# Compilar o binário
RUN CGO_ENABLED=0 go build -o main ./main.go

# Etapa 2: Imagem Final
FROM alpine:latest as runner

# Configurar o diretório de trabalho
WORKDIR /root/

# Copiar o binário da etapa de build
COPY --from=builder /app/main .

COPY --from=builder /app/config /root/config

# Expor a porta usada pelo serviço
EXPOSE 8080

# Comando de execução do binário
CMD ["./main"]