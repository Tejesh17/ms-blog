FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy 

CMD [ "go", "run", "server.go", "comments.go" ]

##  docker build -t tejesh17/blog-ms/comments .
## docker run -t -p 8085:8085 --name blog-ms-comments tejesh17/blog-ms/comments