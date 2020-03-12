FROM golang:latest

RUN mkdir -p /usr/local/go/src/DataApi.Go
WORKDIR /usr/local/go/src/DataApi.Go

ADD . /usr/local/go/src/DataApi.Go
#RUN go get github.com/gin-gonic/gin
#RUN go get -u github.com/jinzhu/gorm
#RUN go mod vendor
RUN go mod download


RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]
