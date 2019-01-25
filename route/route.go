package route

import (
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/user/minio/BucketOperation"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
    Route {
        "GetBucketList",
        "GET",
        "/getbucketlist",
        BucketOperation.GetBucketList,
    },
    
    Route {
        "CreateBucket",
        "POST",
        "/makebucket",
        BucketOperation.CreateBucket,
	},

	Route {
        "RemoveBucket",
        "DELETE",
        "/removebucket",
        BucketOperation.RemoveBucket,
	},

    Route {
        "BucketExist",
        "POST",
        "/bucketexists",
        BucketOperation.BucketExist,
    },  
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes { 
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc
        
        router.
         Methods(route.Method).
         Path(route.Pattern).
         Name(route.Name).
         Handler(handler)
    }
    return router
}