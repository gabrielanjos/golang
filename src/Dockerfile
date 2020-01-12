FROM golang:1.9

WORKDIR /go/src/app

COPY ./src/ ./

RUN apt-get update && go get -d -v github.com/jinzhu/gorm && go get -d -v github.com/gabrielanjos/golang/src  && go build ./main.go
CMD ["main"]