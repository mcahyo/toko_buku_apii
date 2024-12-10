package controllers

import (
	"encoding/json"
	"net/http"
	"toko_buku_api/config"
	"toko_buku_api/models"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUlasan handles POST requests to create a new Ulasan
func CreateUlasan(w http.ResponseWriter, r *http.Request) {
	var ulasan models.Ulasan
	err := json.NewDecoder(r.Body).Decode(&ulasan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Ulasan (Id_Pembeli, Id_Buku, Rating, Ulasan) VALUES (?, ?, ?, ?)`
	_, err = config.DB.Exec(query, ulasan.Id_Pembeli, ulasan.Id_Buku, ulasan.Rating, ulasan.Ulasan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ulasan)
}
