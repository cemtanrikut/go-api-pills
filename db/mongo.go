package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx context.Context
var collection *mongo.Collection

func MongoClient(collection_name string) (context.Context, *mongo.Client, *mongo.Collection) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:LCtfPjhpm1am7HRd@sandbox.0sac2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection = client.Database("datcod-db-test").Collection(collection_name)

	return ctx, client, collection
}
