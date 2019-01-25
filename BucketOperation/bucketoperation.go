package BucketOperation

import (
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Bucket struct 
{
	Name    string `json:"name"`
	Region  string `json:"region"`
	Prefix  string `json:"prefix"`
}


// ********** GetBucketList **********
func GetBucketList(w http.ResponseWriter, _ *http.Request) {
	
	var endpoint = os.Getenv("endpoint")
	var accessKeyID = os.Getenv("accessKeyID")
	var secretAccessKey = os.Getenv("secretAccessKey")
	useSSL := true

  // Initialize minio client object.
  minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
  if err != nil {
	  log.Fatalln(err)
  }
  
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


	var endpoint = os.Getenv("endpoint")
	var accessKeyID = os.Getenv("accessKeyID")
	var secretAccessKey = os.Getenv("secretAccessKey")
	useSSL := true

  // Initialize minio client object.
  minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
  if err != nil {
	  log.Fatalln(err)
  }
  
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
 
  err = minioClient.MakeBucket(bucket.Name ,bucket.Region )
	if err != nil {
    	fmt.Println(err)
    	return
	}

	bytes, err := json.Marshal("Bucket created successfully")
  	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
   writeJsonResponse(w, bytes)
}

// ********** RemoveBucket **********
func RemoveBucket(w http.ResponseWriter, r *http.Request) {

	var endpoint = os.Getenv("endpoint")
	var accessKeyID = os.Getenv("accessKeyID")
	var secretAccessKey = os.Getenv("secretAccessKey")
	useSSL := true

	name := os.Getenv("name")

  // Initialize minio client object.
  minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
  if err != nil {
	  log.Fatalln(err)
  }
  
  err = minioClient.RemoveBucket(name)
  if err != nil {
	bytes, err := json.Marshal("Bucket removed failed.")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJsonResponse(w, bytes)
  }else{
	bytes, err := json.Marshal("Bucket removed successfully.")
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  }
	  writeJsonResponse(w, bytes)
  }
}

// ********** BucketExist **********
func BucketExist(w http.ResponseWriter, r *http.Request) {

	var endpoint = os.Getenv("endpoint")
	var accessKeyID = os.Getenv("accessKeyID")
	var secretAccessKey = os.Getenv("secretAccessKey")
	useSSL := true

  	// Initialize minio client object.
  	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
  	if err != nil {
		  log.Fatalln(err)
  	}

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
		fmt.Println("Bucket found")
		bytes, err := json.Marshal("Bucket found")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		writeJsonResponse(w, bytes)
	}else{
		bytes, err := json.Marshal("Bucket not found")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		writeJsonResponse(w, bytes)
	}
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}