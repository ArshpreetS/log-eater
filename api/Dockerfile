FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o golang-app

CMD [./golang-app]