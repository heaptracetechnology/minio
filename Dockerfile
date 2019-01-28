FROM golang

RUN go get -u github.com/minio/minio-go

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/user/minio

ADD . /go/src/github.com/user/minio

RUN go install github.com/user/minio

ENTRYPOINT minio

EXPOSE 5000
