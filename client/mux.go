package client

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/cemtanrikut/go-api-pills/db"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var ctx context.Context
var collection *mongo.Collection
var router *mux.Router

func MuxHandler() {
	router, ctx, client, collection = db.MongoClient("pill-collection")

	router.HandleFunc("/api/pill/", getPills).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func getPills(w http.ResponseWriter, r *http.Request) {

}
