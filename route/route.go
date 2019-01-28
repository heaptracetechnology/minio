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
        "PUT",
        "/removebucket",
        BucketOperation.RemoveBucket,
	},

    Route {
        "BucketExist",
        "POST",
        "/bucketexists",
        BucketOperation.BucketExist,
    }, 

    Route {
        "GetObjectList",
        "POST",
        "/getobjectlist",
        BucketOperation.GetObjectList,
    },

    Route {
        "GetObject",
        "POST",
        "/getobject",
        BucketOperation.GetObject,
    },
    
    Route {
        "PutObject",
        "PUT",
        "/putobject",
        BucketOperation.PutObject,
    },
    
    Route {
        "CopyObject",
        "POST",
        "/copyobject",
        BucketOperation.CopyObject,
    },
    
    Route {
        "RemoveObject",
        "PUT",
        "/removeobject",
        BucketOperation.RemoveObject,
    },
    
    Route {
        "StatObject",
        "POST",
        "/statobject",
        BucketOperation.StatObject,
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