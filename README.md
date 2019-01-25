# Minio
An OMG service for Minio.

### Installation
* Install minio client
* Install golang version 1.11+

### OMG CLI Command

##### GetBucket

* *omg exec bucketlist -e endpoint=<ENDPOINT> -e accessKeyID=<ACCESS_KEY_ID> -e secretAccessKey=<SECRET_ACCESS_KEY>*

##### CreateBucket

* *omg exec makebucket  -a name=<BUCKET_NAME> -a location=<REGION>  -e endpoint=<ENDPOINT> -e accessKeyID=<ACCESS_KEY_ID> -e secretAccessKey=<SECRET_ACCESS_KEY>*

##### bucketexists

* *omg exec bucketexists  -a name=<BUCKET_NAME> -e endpoint=<ENDPOINT> -e accessKeyID=<ACCESS_KEY_ID> -e secretAccessKey=<SECRET_ACCESS_KEY>*

##### removebucket

* *omg exec removebucket  -a name=<BUCKET_NAME> -e endpoint=<ENDPOINT> -e accessKeyID=<ACCESS_KEY_ID> -e secretAccessKey=<SECRET_ACCESS_KEY>*


### License
[MIT](https://choosealicense.com/licenses/mit/)

