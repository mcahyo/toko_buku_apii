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

// CreateTransaksi handles POST requests to create a new Transaksi
func CreateTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksi models.Transaksi
	err := json.NewDecoder(r.Body).Decode(&transaksi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Transaksi (Id_Pembeli, Total_Harga, Tanggal_Transaksi, Status) VALUES (?, ?, ?, ?)`
	_, err = config.DB.Exec(query, transaksi.Id_Pembeli, transaksi.Total_Harga, transaksi.Tanggal_Transaksi, transaksi.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaksi)
}
