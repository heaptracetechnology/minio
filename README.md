# Minio
An OMG service for Minio.

### Installation
* Install minio client
* Install golang version 1.11+

### OMG CLI Command

##### GetBucket

* *omg exec bucketlist -e endpoint=<ENTER_ENDPOINT> -e accessKeyID=<ACCESS_KEY_ID> -e secretAccessKey=<SECRET_ACCESS_KEY>*

##### CreateBucket

* *omg exec makebucket  -a name=<BUCKET_NAME> -a location=<ENTER_REGION>  -e endpoint=<ENTER_ENDPOINT> -e accessKeyID=<ACCESS_KEY_ID> -e secretAccessKey=<SECRET_ACCESS_KEY>*

##### BucketExists

* *omg exec bucketexists  -a name=<BUCKET_NAME> -e endpoint=<ENTER_ENDPOINT> -e accessKeyID=<ACCESS_KEY_ID> -e secretAccessKey=<SECRET_ACCESS_KEY>*

##### RemoveBucket

* *omg exec removebucket  -a name=<BUCKET_NAME> -e endpoint=<ENTER_ENDPOINT> -e accessKeyID=<ACCESS_KEY_ID> -e secretAccessKey=<SECRET_ACCESS_KEY>*


##### GetObjectList

* *omg exec getobjectlist  -a name=<BUCKET_NAME> -a objectprefix=<OBJECT_PREFIX> -a recursive=<TRUE/FALSE> -e endpoint="play.minio.io:9000" -e accessKeyID="Q3AM3UQ867SPQQA43P2F" -e secretAccessKey="zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"*
* Recursive = "true indicates recursive style listing and false indicates directory style listing delimited by '/'."

##### GetObject

* *omg exec getobject  -a name=<BUCKET_NAME> -a objectname=<OBJECT_NAME> -e endpoint="play.minio.io:9000" -e accessKeyID="Q3AM3UQ867SPQQA43P2F" -e secretAccessKey="zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"*

##### PutObject

* *omg exec putobject  -a name=<BUCKET_NAME> -a objectname=<OBJECT_NAME> -e endpoint="play.minio.io:9000" -e accessKeyID="Q3AM3UQ867SPQQA43P2F" -e secretAccessKey="zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"*

##### CopyObject

* *omg exec copyobject  -a srcbucketname=<SRC_BUCKET_NAME> -a srcobjectname=<SRC_OBJECT_NAME> -a dstbucketname=<DST_BUCKET_NAME> -a dstobjectname=<DST_OBJECT_NAME> -e endpoint="play.minio.io:9000" -e accessKeyID="Q3AM3UQ867SPQQA43P2F" -e secretAccessKey="zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"*

##### RemoveObject

* *omg exec removeobject  -a name=<BUCKET_NAME> -a objectname=<OBJECT_NAME> -e endpoint="play.minio.io:9000" -e accessKeyID="Q3AM3UQ867SPQQA43P2F" -e secretAccessKey="zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"*

##### StatObject

* *omg exec statobject  -a name=<BUCKET_NAME> -a objectname=<OBJECT_NAME> -e endpoint="play.minio.io:9000" -e accessKeyID="Q3AM3UQ867SPQQA43P2F" -e secretAccessKey="zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"*


### License
[MIT](https://choosealicense.com/licenses/mit/)
