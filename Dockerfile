FROM golang:latest

RUN mkdir -p /usr/local/go/src/DataApi.Go /etc/supervisor/conf.d
WORKDIR /usr/local/go/src/DataApi.Go
ADD . /usr/local/go/src/DataApi.Go

RUN go mod download
RUN go build ./main.go

RUN \
    apt-get update && apt-get install --assume-yes apt-utils ; \
    apt-get -y install python python-dev python-setuptools python-pip

RUN pip install supervisor==3.2

COPY ./supervisor/supervisord.conf /etc/
COPY ./supervisor/supervisord-gin.conf /etc/supervisor/conf.d/

EXPOSE 8080
CMD ["./main"]
