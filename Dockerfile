#TODO :: Need to implement 

FROM golang

RUN go get -u github.com/minio/minio-go

WORKDIR Minio

ADD . Minio

ENTRYPOINT Minio RUN

EXPOSE 5000
