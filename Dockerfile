FROM golang:latest

RUN apt-get update && \
    apt-get -y install vim

WORKDIR /go/src/gqlgen-todos
RUN go mod init gqlgen-todos
# mongoのモジュール
RUN go get gopkg.in/mgo.v2
# gqlgenのモジュール
RUN go get github.com/99designs/gqlgen
# gqlgenの枠組み
RUN go run github.com/99designs/gqlgen init
ADD . /go/src/gqlgen-todos