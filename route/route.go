package route

import (
    "github.com/gorilla/mux"
    "github.com/user/minio/BucketOperation"
    "github.com/user/minio/ObjectOperation"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "GetBucketList",
        "GET",
        "/getbucketlist",
        BucketOperation.GetBucketList,
    },

    Route{
        "CreateBucket",
        "POST",
        "/makebucket",
        BucketOperation.CreateBucket,
    },

    Route{
        "BucketExist",
        "POST",
        "/bucketexists",
        BucketOperation.BucketExist,
    },

    Route{
        "GetObjectList",
        "POST",
        "/getobjectlist",
        ObjectOperation.GetObjectList,
    },

    Route{
        "GetObject",
        "POST",
        "/getobject",
        ObjectOperation.GetObject,
    },

    Route{
        "PutObject",
        "PUT",
        "/putobject",
        ObjectOperation.PutObject,
    },

    Route{
        "CopyObject",
        "POST",
        "/copyobject",
        ObjectOperation.CopyObject,
    },

    Route{
        "RemoveObject",
        "PUT",
        "/removeobject",
        ObjectOperation.RemoveObject,
    },

    Route{
        "StatObject",
        "POST",
        "/statobject",
        ObjectOperation.StatObject,
    },

    Route{
        "GetBucketPolicy",
        "POST",
        "/getbucketpolicy",
        BucketOperation.GetBucketPolicy,
    },

    Route{
        "FPutObject",
        "PUT",
        "/fputobject",
        ObjectOperation.FPutObject,
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
