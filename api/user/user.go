package user

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserData struct {
	name      string
	email     string
	password  string
	gender    string
	birthdate time.Time
	active    bool
}

func SignUp(resp http.ResponseWriter, req *http.Request, client *mongo.Client, collection *mongo.Collection) (UserData, error) {
	resp.Header().Set("Content-Type", "application/json")
	var user UserData
	json.NewDecoder(req.Body).Decode(&user)
	user.password = base64.StdEncoding.EncodeToString([]byte(user.password))
	user.active = false

	checkEmail := CheckEmail(user.email, client, collection)
	if checkEmail {
		return UserData{}, fmt.Errorf("%s", "This mail address already exist")
	}

	_, insertErr := collection.InsertOne(context.Background(), user)
	if insertErr != nil {
		return UserData{}, insertErr
	}

	return user, nil

}

func LogIn(resp http.ResponseWriter, req *http.Request, client *mongo.Client, ctx context.Context, collection *mongo.Collection) error {
	resp.Header().Set("Content-Type", "application/json")
	var user, dbUser UserData

	json.NewDecoder(req.Body).Decode(&user)

	err := collection.FindOne(context.Background(), bson.M{"email": user.email, "active": true}).Decode(&dbUser)

	if err != nil {
		return err

	}

	user.password = base64.StdEncoding.EncodeToString([]byte(user.password))

	userPass := []byte(user.password)
	dbPass := []byte(dbUser.password)

	fmt.Println(userPass, dbPass)
	fmt.Println(user.password, dbUser.password)

	//passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	res := bytes.Equal(userPass, dbPass)
	if !res {
		return fmt.Errorf("%s", "Invalid Password")
	}

	return nil

}

func CheckEmail(email string, client *mongo.Client, collection *mongo.Collection) bool {
	var dbUser UserData
	err := collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&dbUser)
	fmt.Println("data - ", err)
	if err == nil {
		fmt.Println(email, " already exist")
		return true
	}
	return false
}
