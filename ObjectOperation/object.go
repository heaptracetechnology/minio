package ObjectOperation

import (
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type BucketObject struct {
	Name          string `json:"name"`
	Prefix        string `json:"prefix"`
	ObjectName    string `json:"objectname"`
	FileName      string `json:"filename"`
	FilePath      string `json:"filepath"`
	SrcBucketName string `json:"srcbucketname"`
	SrcObjectName string `json:"srcobjectname"`
	DstBucketName string `json:"dstbucketname"`
	DstObjectName string `json:"dstobjectname"`
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

// ********** GetObjectList **********
func GetObjectList(w http.ResponseWriter, r *http.Request) {

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close() //defer => work as finally block(ensure that a function call is performed later in a program’s execution)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucketobj BucketObject
	err = json.Unmarshal(b, &bucketobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)
	isRecursive := true

	objectCh := minioClient.ListObjects(bucketobj.Name, bucketobj.Prefix, isRecursive, doneCh)
	bytes, _ := json.Marshal(objectCh)
	writeJsonResponse(w, bytes)
}

// ********** GetObject **********
func GetObject(w http.ResponseWriter, r *http.Request) {

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucketobj BucketObject
	err = json.Unmarshal(b, &bucketobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	object, err := minioClient.GetObject(bucketobj.Name, bucketobj.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	localFile, err := os.Create("newfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(localFile)
	bytes, _ := json.Marshal(object)
	writeJsonResponse(w, bytes)

}

//********** PutObject **********
func PutObject(w http.ResponseWriter, r *http.Request) {

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucketobj BucketObject
	err = json.Unmarshal(b, &bucketobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("------bucket--->>>>", bucketobj)
	file, err := os.Open(bucketobj.FileName)
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

	n, err := minioClient.PutObject(bucketobj.Name, bucketobj.ObjectName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	bytes, _ := json.Marshal(n)
	writeJsonResponse(w, bytes)
}

//********** FPutObject **********
func FPutObject(w http.ResponseWriter, r *http.Request) {

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucketobj BucketObject
	err = json.Unmarshal(b, &bucketobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	n, err := minioClient.FPutObject(bucketobj.Name, bucketobj.ObjectName, bucketobj.FilePath, minio.PutObjectOptions{ContentType: "text/plain"})
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

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucketobj BucketObject
	err = json.Unmarshal(b, &bucketobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Source object
	src := minio.NewSourceInfo(bucketobj.SrcBucketName, bucketobj.SrcObjectName, nil)

	// Destination object
	dst, err := minio.NewDestinationInfo(bucketobj.DstBucketName, bucketobj.DstObjectName, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Copy object call
	err = minioClient.CopyObject(dst, src)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		bytes, _ := json.Marshal("Copy successful")
		writeJsonResponse(w, bytes)
	}
}

//********** RemoveObject **********
func RemoveObject(w http.ResponseWriter, r *http.Request) {

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucketobj BucketObject
	err = json.Unmarshal(b, &bucketobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = minioClient.RemoveObject(bucketobj.Name, bucketobj.ObjectName)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		bytes, _ := json.Marshal("Object removed successfully")
		writeJsonResponse(w, bytes)
	}
}

//********** StatObject **********
func StatObject(w http.ResponseWriter, r *http.Request) {

	minioClient, err := MinioClient()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var bucketobj BucketObject
	err = json.Unmarshal(b, &bucketobj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	objInfo, err := minioClient.StatObject(bucketobj.Name, bucketobj.ObjectName, minio.StatObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		bytes, _ := json.Marshal(objInfo)
		writeJsonResponse(w, bytes)
	}
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}
