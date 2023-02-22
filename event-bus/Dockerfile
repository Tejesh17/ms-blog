FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]

##  docker build -t tejesh17/blog-ms/event-bus .
## docker run -t -p 8085:8085 --name blog-ms-event-bus tejesh17/blog-ms/event-bus