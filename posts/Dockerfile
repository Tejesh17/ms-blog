FROM golang:1.20-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]

## docker build -t tejesh17/blog-ms-posts . 
## docker run --name blog-ms-posts -t -p 8080:8080 tejesh17/blog-ms-posts