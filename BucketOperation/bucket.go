package BucketOperation

import (
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/minio/result"
	"github.com/minio/minio-go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Bucket struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Prefix string `json:"prefix"`
}

type Message struct {
    Success string `json:"success"`
    Message string `json:"message"`
}

// ********** MinioClient **********
func MinioClient() (*minio.Client, error) {

	var endpoint = os.Getenv("END_POINT")
	var accessKeyID = os.Getenv("ACCESS_KEY_ID")
	var secretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient, nil
}

// ********** GetBucketList **********
func GetBucketList(w http.ResponseWriter, _ *http.Request) {

	minioClient, err := MinioClient()

	buckets, err := minioClient.ListBuckets()
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
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

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = minioClient.MakeBucket(bucket.Name, bucket.Region)
	if err != nil {
		exists, err := minioClient.BucketExists(bucket.Name)
		if err == nil && exists {
			msg := Message{"false", "We already own " + bucket.Name}
			msgbytes, _ := json.Marshal(msg)
			writeJsonResponse(w, msgbytes)
			return
		} else {
			log.Fatalln(err)
		}
	} else {
		msg := Message{"true", "Successfully created " + bucket.Name}
		msgbytes, _ := json.Marshal(msg)
		writeJsonResponse(w, msgbytes)
		return
	}
}

// ********** BucketExist **********
func BucketExist(w http.ResponseWriter, r *http.Request) {

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	found, err := minioClient.BucketExists(bucket.Name)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, _ := json.Marshal(msg)
		writeJsonResponse(w, msgbytes)
		return
	}
	if found {
		msg := Message{"true", "Bucket found"}
		msgbytes, _ := json.Marshal(msg)
		writeJsonResponse(w, msgbytes)
		return
	} else {
		msg := Message{"false", "Bucket not found"}
		msgbytes, _ := json.Marshal(msg)
		writeJsonResponse(w, msgbytes)
		return
	}
}

// ********** GetBucketPolicy **********
func GetBucketPolicy(w http.ResponseWriter, r *http.Request) {

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	policy, err := minioClient.GetBucketPolicy(bucket.Name)
	if err != nil {
		msg := Message{"false", "Bucket policy not found"}
		msgbytes, _ := json.Marshal(msg)
		writeJsonResponse(w, msgbytes)
		return
	} else {
		bytes, _ := json.Marshal(policy)
		writeJsonResponse(w, bytes)
	}
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
