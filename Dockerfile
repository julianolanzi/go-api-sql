# Use a imagem oficial do Golang como base para o build
FROM golang:1.22.1 AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copiar os arquivos de dependências
COPY go.mod go.sum ./

# Baixar as dependências do Go
RUN go mod download

# Copiar o código-fonte para o contêiner
COPY . .

# Verificar se o arquivo main.go está no local correto
RUN ls -la /go/src/app/cmd

# Construir o aplicativo Go
RUN go build -o main ./cmd/main.go

# Usar uma imagem mínima para a execução
FROM alpine:latest

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copiar o binário construído da fase anterior
COPY --from=builder /go/src/app/main .

# Expor a porta que o aplicativo irá usar
EXPOSE 1111

# Executar o binário
CMD ["./main"]