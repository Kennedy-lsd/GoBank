package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func idParser(w http.ResponseWriter, r *http.Request) (int64, error) {
	pathSegments := strings.Split(r.URL.Path, "/")

	if len(pathSegments) < 3 {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return 0, nil
	}

	idStr := pathSegments[3]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("Error converting ID:", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return 0, err
	}

	return id, nil
}

func writeJSON(w http.ResponseWriter, status int, resp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}
