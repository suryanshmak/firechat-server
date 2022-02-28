FROM golang:1.18rc1-alpine3.15

WORKDIR /app

COPY go.mod ./

RUN go mod tidy 

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/web/*"]
