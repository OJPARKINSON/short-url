package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ojparkinson/shortUrl/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type shortenReq struct {
	URL string `json:"url"`
}

type url struct {
	Id        string
	Url       string
	ShortCode string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ShortenHandler struct{}

func (s *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		collection, dberr := db.ConnectToCollection()
		if dberr != nil {
			fmt.Println("failed to connect to the db: ", dberr)
		}

		var reqBody shortenReq
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			fmt.Println("Error decoding body: ", err)
			w.WriteHeader(400)
		}

		url := url{Url: reqBody.URL, ShortCode: "abc123", CreatedAt: time.Now(), UpdatedAt: time.Now()}

		result, insertErr := collection.InsertOne(context.TODO(), url)
		if insertErr != nil {
			fmt.Println("Error inserting url: ", err)
			w.WriteHeader(500)
		} else if result.Acknowledged == false {
			fmt.Println("insert not acknowledged: ", err)
			w.WriteHeader(500)
		}

		url.Id = result.InsertedID.(bson.ObjectID).Hex()

		fmt.Println(result.Acknowledged, result.InsertedID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(url)
	}

	w.WriteHeader(404)
}
