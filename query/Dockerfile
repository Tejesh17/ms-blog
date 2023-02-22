FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy 

CMD [ "go", "run", "server.go", "query.go" ]

##  docker build -t tejesh17/blog-ms/query .
## docker run -t -p 8085:8085 --name blog-ms-query tejesh17/blog-ms/query