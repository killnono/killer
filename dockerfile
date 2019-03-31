FROM golang:latest

WORKDIR $GOPATH/src/app
COPY . $GOPATH/src/app

RUN go build -o main . 
EXPOSE 12345
ENTRYPOINT ["./main"]
