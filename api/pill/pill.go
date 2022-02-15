package pill

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type pill struct {
	name              string `json:"name"`
	barcode           string `json:"barcode"`
	atc_code          string `json:"atc_code"`
	atc_name          string `json:"atc_name"`
	company_name      string `json:"company_name"`
	prescription_type string `json:"prescription_type"`
	status            bool   `json:"status"`
}

func getByName(resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getByBarcode(resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getAll(resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) ([]byte, error) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var pillMList []primitive.M

	cursor, err := collection.Find(context.Background(), bson.M{"status": true})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var pll bson.M
		if err = cursor.Decode(&pll); err != nil {
			return nil, err
		}
		pillMList = append(pillMList, pll)
	}
	defer cursor.Close(context.Background())

	jsonResult, err := json.Marshal(pillMList)
	if err != nil {
		return nil, err
	}

	return jsonResult, nil

}
