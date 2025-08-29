package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
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
		CreateShorten(w, r)
		return
	}

	w.WriteHeader(404)
}

func CreateShorten(w http.ResponseWriter, r *http.Request) {
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

	shortCode := generateShortCode()
	url := url{Url: reqBody.URL, ShortCode: shortCode, CreatedAt: time.Now(), UpdatedAt: time.Now()}

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

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}

	return string(result)
}
