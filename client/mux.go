package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	apiPill "github.com/cemtanrikut/go-api-pills/api"
	"github.com/cemtanrikut/go-api-pills/db"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var ctx context.Context
var collection *mongo.Collection
var router *mux.Router

func MuxHandler() {
	router, ctx, client, collection = db.MongoClient("pills-collection")

	router.HandleFunc("/api/pill/{barcode}", getPill).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func getPill(w http.ResponseWriter, r *http.Request) {
	barcode := mux.Vars(r)["barcode"]
	result := apiPill.getByBarcode(barcode, w, r, collection)
	jsonResult, jsonError := json.Marshal(result)
	if jsonError != nil {
	}
	w.Write(jsonResult)
}
