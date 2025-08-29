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
	Url       string    `bson:"url,omitempty"`
	ShortCode string    `bson:"shortcode,omitempty"`
	CreatedAt time.Time `bson:"createdat,omitempty"`
	UpdatedAt time.Time `bson:"updatedat,omitempty"`
}

type ShortenHandler struct{}

func (s *ShortenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.GetShorten(w, r)
		return
	case http.MethodPost:
		s.CreateShorten(w, r)
		return
	}

	w.WriteHeader(404)
}

func (s *ShortenHandler) CreateShorten(w http.ResponseWriter, r *http.Request) {
	var reqBody shortenReq
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println("Error decoding body: ", err)
		w.WriteHeader(400)
		return
	}

	collection, dberr := db.ConnectToCollection()
	if dberr != nil {
		fmt.Println("failed to connect to the db: ", dberr)
		return
	}

	shortCode := generateShortCode()
	url := url{Url: reqBody.URL, ShortCode: shortCode, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	result, insertErr := collection.InsertOne(context.TODO(), url)
	if insertErr != nil {
		fmt.Println("Error inserting url: ", err)
		w.WriteHeader(500)
	} else if !result.Acknowledged {
		fmt.Println("insert not acknowledged: ", err)
		w.WriteHeader(500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(url)
}

func (s *ShortenHandler) GetShorten(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("shortcode")
	if shortCode == "" {
		http.Error(w, "Short code required", http.StatusBadRequest)
		return
	}

	collection, dberr := db.ConnectToCollection()
	if dberr != nil {
		http.Error(w, "Failed to connect to the DB", http.StatusInternalServerError)
		return
	}
	filter := bson.D{{Key: "shortcode", Value: shortCode}}

	var reqBody url
	result := collection.FindOne(context.TODO(), filter).Decode(&reqBody)
	if result != nil {
		w.WriteHeader(400)
		http.Error(w, "Failed to find the short url", http.StatusNotFound)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(reqBody)
}

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, 10)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}

	return string(result)
}
