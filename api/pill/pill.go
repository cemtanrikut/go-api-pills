package pill

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type pill struct {
	name              string `json:"name"`
	barcode           string `json:"barcode"`
	atc_code          string `json:"atc_code"`
	atc_name          string `json:"atc_name"`
	company_name      string `json:"company_name"`
	prescription_type string `json:"prescription_type"`
	status            string `json:"status"`
}

func getByName(resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getByBarcode(resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) {

}

func getAll(resp http.ResponseWriter, req *http.Request, collection *mongo.Collection) {

}
