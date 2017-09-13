FROM golang:1.8

RUN mkdir -p /go/src/github.com/fesiqp/jwtauth/api
WORKDIR /go/src/github.com/fesiqp/jwtauth/api 

COPY . /go/src/github.com/fesiqp/jwtauth/api

RUN go get github.com/codegangsta/gin
RUN go get -u github.com/jinzhu/gorm
RUN go-wrapper download
RUN go-wrapper install

CMD gin run