# Etapa 1: Build
FROM golang:1.23 AS builder

# Configurar o diretório de trabalho no container
WORKDIR /app

# Copiar os arquivos do projeto para o container
COPY go.mod ./
RUN go mod tidy

# Copiar o restante do código
COPY . .

# Compilar o binário
RUN CGO_ENABLED=0 go build -o main ./cmd/main.go

RUN curl -o /app/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh \
    && chmod +x /app/wait-for-it.sh

# Etapa 2: Imagem Final
FROM alpine:latest as runner

# Instalar o shell (se necessário) ou o bash
RUN apk update && apk add --no-cache bash

# Configurar o diretório de trabalho
WORKDIR /root/

# Copiar o binário da etapa de build
COPY --from=builder /app/main .

COPY --from=builder /app/configs /root/configs

# Copiar o script wait-for-it.sh para o contêiner final
COPY --from=builder /app/wait-for-it.sh /root/wait-for-it.sh

# Expor a porta usada pelo serviço
EXPOSE 8080

# Comando de execução do binário
CMD ["./wait-for-it.sh", "db:3306", "--", "./main"]