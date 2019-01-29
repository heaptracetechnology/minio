package BucketOperation

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"io"
)

type Bucket struct 
{
	Name    	  string `json:"name"`
	Region  	  string `json:"region"`
	Prefix  	  string `json:"prefix"`
	ObjectName 	  string `json:"objectname"`
	FileName 	  string `json:"filename"`
	FilePath      string `json:"filepath"`
	SrcBucketName string `json:"srcbucketname"`
	SrcObjectName string `json:"srcobjectname"`
	DstBucketName string `json:"dstbucketname"`
	DstObjectName string `json:"dstobjectname"`
}

// ********** MinioClient **********
func MinioClient() (*minio.Client, error){

	var endpoint = os.Getenv("endpoint")
	var accessKeyID = os.Getenv("accessKeyID")
	var secretAccessKey = os.Getenv("secretAccessKey")
	useSSL := true

  	// Initialize minio client object.
  	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
  	if err != nil {
		log.Fatalln(err)
	  }
	return minioClient,nil
}

// ********** GetBucketList **********
func GetBucketList(w http.ResponseWriter, _ *http.Request) {
		
	minioClient,err := MinioClient()
  
  	buckets, err := minioClient.ListBuckets()
  	if err != nil {
	  fmt.Println(err)
	  return
  	}

	bytes, err := json.Marshal(buckets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
  	writeJsonResponse(w, bytes)
}

// ********** CreateBucket **********
func CreateBucket(w http.ResponseWriter, r *http.Request) {

	minioClient,err := MinioClient()
  
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = minioClient.MakeBucket(bucket.Name ,bucket.Region)
    if err != nil {
        exists, err := minioClient.BucketExists(bucket.Name)
        if err == nil && exists {
			result := "We already own " + bucket.Name
			bytes, err := json.Marshal(result)
	  		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
	   		writeJsonResponse(w, bytes)
        } else {
            log.Fatalln(err)
        }
    } else {
		result := "Successfully created " + bucket.Name
		bytes, err := json.Marshal(result)
	  	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	   	writeJsonResponse(w, bytes)
    }
}

// ********** RemoveBucket **********
func RemoveBucket(w http.ResponseWriter, r *http.Request) {

	minioClient,err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
  	err = minioClient.RemoveBucket(bucket.Name)
  	if err != nil {
		bytes, _ := json.Marshal("Bucket removed failed.")
		writeJsonResponse(w, bytes)
    }else{
		bytes, _ := json.Marshal("Bucket removed successfully.")
		writeJsonResponse(w, bytes)
    }
}

// ********** BucketExist **********
func BucketExist(w http.ResponseWriter, r *http.Request) {

	minioClient,err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	found, err := minioClient.BucketExists(bucket.Name)
	if err != nil {
    	fmt.Println(err)
    	return
	}
	if found {
		bytes, _ := json.Marshal("Bucket found")
		writeJsonResponse(w, bytes)
	}else{
		bytes, _ := json.Marshal("Bucket not found")
		writeJsonResponse(w, bytes)
	}
}

// ********** GetObjectList **********
func GetObjectList(w http.ResponseWriter, r *http.Request) {
	
	minioClient,err := MinioClient()
 
   	b, err := ioutil.ReadAll(r.Body)
   	defer r.Body.Close() //defer => work as finally block(ensure that a function call is performed later in a program’s execution)
   	if err != nil {
	   http.Error(w, err.Error(), 500)
	   return
   	}

   	// Unmarshal
   	var bucket Bucket
   	err = json.Unmarshal(b, &bucket)
   	if err != nil {
	   http.Error(w, err.Error(), 500)
	   return
   	}
    doneCh := make(chan struct{})

    // Indicate to our routine to exit cleanly upon return.
    defer close(doneCh)
    isRecursive := true

    objectCh := minioClient.ListObjects(bucket.Name, bucket.Prefix, isRecursive, doneCh)
    bytes, _ := json.Marshal(objectCh)
	writeJsonResponse(w, bytes)
}

// ********** GetObject **********
func GetObject(w http.ResponseWriter, r *http.Request) {

	minioClient,err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	object, err := minioClient.GetObject(bucket.Name, bucket.ObjectName, minio.GetObjectOptions{})
	if err != nil {
    	fmt.Println(err)
    	return
	}

	localFile, err := os.Create(bucket.FilePath)
	if err != nil {
    	fmt.Println(err)
    	return
	}

	if _, err = io.Copy(localFile, object); err != nil {
    	fmt.Println(err)
    	return
	}

	bytes, _ := json.Marshal(object)
    writeJsonResponse(w, bytes)

}

//********** PutObject **********
func PutObject(w http.ResponseWriter, r *http.Request) {

	minioClient,err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	file, err := os.Open(bucket.FileName)
	if err != nil {
    	fmt.Println(err)
    	return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
    	fmt.Println(err)
    	return
	}

	n, err := minioClient.PutObject(bucket.Name, bucket.ObjectName, file, fileStat.Size(), minio.PutObjectOptions{ContentType:"application/octet-stream"})
	if err != nil {
	    fmt.Println(err)
    	return
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	bytes, _ := json.Marshal(n)
    writeJsonResponse(w, bytes)
}

//********** CopyObject **********
func CopyObject(w http.ResponseWriter, r *http.Request) {

	minioClient,err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Source object
	src := minio.NewSourceInfo(bucket.SrcBucketName, bucket.SrcObjectName, nil)

	// Destination object
	dst, err := minio.NewDestinationInfo(bucket.DstBucketName, bucket.DstObjectName, nil, nil)
	if err != nil {
    	fmt.Println(err)
    	return
	}

	// Copy object call
	err = minioClient.CopyObject(dst, src)
	if err != nil {
    	fmt.Println(err)
    	return
	}else{
		bytes, _ := json.Marshal("Copy successful")
   		writeJsonResponse(w, bytes)
	}
}

//********** RemoveObject **********
func RemoveObject(w http.ResponseWriter, r *http.Request) {

	minioClient,err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = minioClient.RemoveObject(bucket.Name, bucket.ObjectName)
	if err != nil {
    	fmt.Println(err)
    	return
	}else{
		bytes, _ := json.Marshal("Object removed successfully")
   		writeJsonResponse(w, bytes)
	}
}

//********** StatObject **********
func StatObject(w http.ResponseWriter, r *http.Request) {

	minioClient,err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	objInfo, err := minioClient.StatObject(bucket.Name, bucket.ObjectName, minio.StatObjectOptions{})
	if err != nil {
    	fmt.Println(err)
    	return
	}else{
		bytes, _ := json.Marshal(objInfo)
   		writeJsonResponse(w, bytes)
	}
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}