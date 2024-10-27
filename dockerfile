FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Move para o diretório cmd antes de construir
WORKDIR /app/api

# Construa o arquivo main.go dentro do diretório cmd
RUN go build -o /app/main

# Volte para o diretório app
WORKDIR /app

EXPOSE 8080

CMD ./main
