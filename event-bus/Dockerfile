FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod init example.go/event-bus \
    && go mod tidy \
    && go build -o app

CMD [ "./app" ]

##  docker build -t tejesh17/blog-ms/event-bus .
## docker run -t -p 8085:8085 --name blog-ms-event-bus tejesh17/blog-ms/event-bus