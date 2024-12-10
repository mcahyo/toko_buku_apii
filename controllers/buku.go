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

// CreateBuku handles POST requests to create a new Buku
func CreateBuku(w http.ResponseWriter, r *http.Request) {
	var buku models.Buku
	err := json.NewDecoder(r.Body).Decode(&buku)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Buku (Judul, Penulis, Jenis, ISBN, Tahun_Terbit, Harga, Stok) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = config.DB.Exec(query, buku.Judul, buku.Penulis, buku.Jenis, buku.ISBN, buku.Tahun_Terbit, buku.Harga, buku.Stok)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buku)
}

// GetBuku handles GET requests to fetch all Buku
func GetBuku(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM Buku")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bukus []models.Buku
	for rows.Next() {
		var buku models.Buku
		if err := rows.Scan(&buku.Id_Buku, &buku.Judul, &buku.Penulis, &buku.Jenis, &buku.ISBN, &buku.Tahun_Terbit, &buku.Harga, &buku.Stok); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bukus = append(bukus, buku)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bukus)
}

// GetBukuByID handles GET requests to fetch a Buku by ID
func GetBukuByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var buku models.Buku
	query := `SELECT * FROM Buku WHERE Id_Buku = ?`
	err = config.DB.QueryRow(query, id).Scan(&buku.Id_Buku, &buku.Judul, &buku.Penulis, &buku.Jenis, &buku.ISBN, &buku.Tahun_Terbit, &buku.Harga, &buku.Stok)
	if err != nil {
		http.Error(w, "Buku not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buku)
}
