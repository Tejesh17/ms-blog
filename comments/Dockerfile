FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]
##  docker build -t tejesh17/blog-ms-comments .
## docker run -t -p 8081:8081 --name blog-ms-comments tejesh17/blog-ms/comments