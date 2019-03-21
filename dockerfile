# FROM alpine:3.4

# ADD ./kill /go/bin/kill

# WORKDIR /app
# CMD ["/go/bin/kill"]
# EXPOSE 12345




FROM golang:latest
# RUN mkdir /app 
# ADD . /app/ 
# COPY . /app/

WORKDIR $GOPATH/src/app
COPY . $GOPATH/src/app
# RUN go build .


# ADD ./github.com  /app/github.com 
# RUN go get github.com/gin-gonic/gin
# RUN go get github.com/gorilla/websocket
# WORKDIR /app 
# go get
RUN go build -o main . 
EXPOSE 12345
ENTRYPOINT ["./main"]

# CMD ["/app/main"]