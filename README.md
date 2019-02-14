# Minio as a microservice
An OMG service to create, get and get list of bucket and object on Minio.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)

This microservice's goal is to to create, get and get list of bucket and object on Minio

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```

### CLI

##### Bucket Exists
```sh
$ omg run existsbucket  -a name=<*BUCKET_NAME*> -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```
##### Create Bucket
```sh
$ omg run makebucket  -a name=<*BUCKET_NAME*> -a location=<*ENTER_REGION*>  -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```
##### Get Bucket List
```sh
$ omg run listbuckets -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```
##### Get Bucket Object List
```sh
$ omg run listobjects  -a name=<*BUCKET_NAME*> -a objectprefix=<*OBJECT_PREFIX*> -a recursive=<*TRUE/FALSE*> -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```
##### Get Bucket Object
```sh
$ omg run getobject  -a name=<*BUCKET_NAME*> -a objectname=<*OBJECT_NAME*> -a filepath=<*FILE_PATH*> -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```
##### Put Bucket Object
```sh
$ omg run putobject  -a name=<*BUCKET_NAME*> -a objectname=<*OBJECT_NAME*> -e endpoint="play.minio.io:9000" -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```
##### Copy Bucket Object
```sh
$ omg run copyobject  -a srcbucketname=<*SRC_BUCKET_NAME*> -a srcobjectname=<*SRC_OBJECT_NAME*> -a dstbucketname=<*DST_BUCKET_NAME*> -a dstobjectname=<*DST_OBJECT_NAME*> -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```
##### Stat Bucket Object
```sh
$ omg run statobject  -a name=<*BUCKET_NAME*> -a objectname=<*OBJECT_NAME*> -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```
##### Put File Bucket Object
```sh
$ omg run fputobject  -a name=<*BUCKET_NAME*> -a objectname=<*OBJECT_NAME*> -a filepath=<*FILE_PATH*> -e endpoint=<*ENTER_ENDPOINT*> -e accessKeyID=<*ACCESS_KEY_ID*> -e secretAccessKey=<*SECRET_ACCESS_KEY*>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Installation
* Install minio client
* Install golang version 1.11+

## Docker
### Build
```
docker build --rm -f "Dockerfile" -t minio:latest .
```
### RUN
```
docker run -p 5000:5000 minio:latest
```

