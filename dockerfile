FROM golang:latest
RUN mkdir /app 
ADD . /app/ 
# ADD ./github.com  /app/github.com 
RUN go get github.com/gin-gonic/gin
WORKDIR /app 
# go get
RUN go build -o main . 
CMD ["/app/main"]