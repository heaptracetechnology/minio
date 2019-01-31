package BucketOperation

import (
	"encoding/json"
	"fmt"
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

// ********** MinioClient **********
func MinioClient() (*minio.Client, error) {

	var endpoint = os.Getenv("endpoint")
	var accessKeyID = os.Getenv("accessKeyID")
	var secretAccessKey = os.Getenv("secretAccessKey")
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

	minioClient, err := MinioClient()

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

	err = minioClient.MakeBucket(bucket.Name, bucket.Region)
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

	minioClient, err := MinioClient()

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
	} else {
		bytes, _ := json.Marshal("Bucket removed successfully.")
		writeJsonResponse(w, bytes)
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
	} else {
		bytes, _ := json.Marshal("Bucket not found")
		writeJsonResponse(w, bytes)
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
	// Unmarshal
	var bucket Bucket
	err = json.Unmarshal(b, &bucket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	policy, err := minioClient.GetBucketPolicy(bucket.Name)
	if err != nil {
		fmt.Println(err)
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
