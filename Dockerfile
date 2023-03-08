FROM golang:1.17.7-alpine

WORKDIR /app
COPY . .

RUN apk upgrade && apk add git

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", "./.air.toml"]
