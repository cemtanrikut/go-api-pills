package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PillData struct {
	Name              string `json:"name"`
	Barcode           string `json:"barcode"`
	Atc_code          string `json:"atc_code"`
	Atc_name          string `json:"atc_name"`
	Company_name      string `json:"company_name"`
	Prescription_type string `json:"prescription_type"`
	Status            bool   `json:"status"`
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

func GetByBarcode(barcode string, resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) []byte {
	resp.Header().Set("Content-Type", "application/json")

	var data PillData

	pillData := collection.FindOne(context.Background(), bson.M{"barcode": barcode})
	err := pillData.Decode(&data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	jsonResult, _ := json.Marshal(data)

	return jsonResult
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
