package main

import (
	"gopkg.in/mgo.v2"
	"graphql/db"
	"graphql/gq"
	"net/http"
	"runtime"
)

func main() {
	db.Mongo, _ = mgo.Dial("mongodb://127.0.0.1:27017")

	db.Mongo.SetMode(mgo.Monotonic, true)
	db.Mongo.SetPoolLimit(600)
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	http.HandleFunc("/graphql", gq.GraphQlHandler)
	http.ListenAndServe(":8080", nil)
}
