package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toko_buku_api/config"
	"toko_buku_api/models"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePembeli handles POST requests to create a new Pembeli
func CreatePembeli(w http.ResponseWriter, r *http.Request) {
	var pembeli models.Pembeli
	err := json.NewDecoder(r.Body).Decode(&pembeli)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Pembeli (nama_Pembeli, kata_Sandi, Alamat_Email) VALUES (?, ?, ?)`
	_, err = config.DB.Exec(query, pembeli.Nama_Pembeli, pembeli.Kata_Sandi, pembeli.Alamat_Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pembeli)
}

// GetPembeli handles GET requests to fetch all Pembeli
func GetPembeli(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM Pembeli")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pembelis []models.Pembeli
	for rows.Next() {
		var pembeli models.Pembeli
		if err := rows.Scan(&pembeli.Id_Pembeli, &pembeli.Nama_Pembeli, &pembeli.Kata_Sandi, &pembeli.Alamat_Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pembelis = append(pembelis, pembeli)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pembelis)
}

// GetPembeliByID handles GET requests to fetch a Pembeli by ID
func GetPembeliByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var pembeli models.Pembeli
	query := `SELECT * FROM Pembeli WHERE Id_Pembeli = ?`
	err = config.DB.QueryRow(query, id).Scan(&pembeli.Id_Pembeli, &pembeli.Nama_Pembeli, &pembeli.Kata_Sandi, &pembeli.Alamat_Email)
	if err != nil {
		http.Error(w, "Pembeli not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pembeli)
}
