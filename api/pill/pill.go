package api

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PillData struct {
	name              string `json:"name"`
	barcode           string `json:"barcode"`
	atc_code          string `json:"atc_code"`
	atc_name          string `json:"atc_name"`
	company_name      string `json:"company_name"`
	prescription_type string `json:"prescription_type"`
	status            bool   `json:"status"`
}

func getByName(name string, resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) (*PillData, error) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var data PillData

	pillData := collection.FindOne(context.Background(), bson.M{
		"$and": []bson.M{
			{"name": name},
			{"status": true},
		},
	})
	err := pillData.Decode(&data)
	if err != nil {
		return &PillData{}, err
	}

	return &data, nil
}

func getByExistsName(name string, resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) ([]byte, error) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var pillMList []primitive.M

	filter := bson.D{{"name", bson.D{{"$exists", name}}}}
	cursor, err := collection.Find(context.Background(), filter)
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
	//db.collection.find( { "arr": { "$exsits": true } } )
}

func getByBarcode(barcode string, resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) (*PillData, error) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var data PillData

	pillData := collection.FindOne(context.Background(), bson.M{
		"$and": []bson.M{
			{"barcode": barcode},
			{"status": true},
		},
	})
	err := pillData.Decode(&data)
	if err != nil {
		return &PillData{}, err
	}

	return &data, nil
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
