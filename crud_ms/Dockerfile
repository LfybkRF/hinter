FROM golang:1.21.1

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download

EXPOSE 3001