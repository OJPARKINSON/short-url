package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ojparkinson/shortUrl/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StatsHandler struct{}

func (s *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.GetStats(w, r)
		return
	}

	w.WriteHeader(404)
}

func (s *StatsHandler) GetStats(w http.ResponseWriter, r *http.Request) {
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

	var resBody Url
	result := collection.FindOne(context.TODO(), filter).Decode(&resBody)
	if result != nil {
		if result == mongo.ErrNoDocuments {
			http.Error(w, "Failed to find the short url", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to get the document", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resBody)
}
