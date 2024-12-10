package controllers

import (
	"encoding/json"
	"net/http"
	"toko_buku_api/config"
	"toko_buku_api/models"
)

// CreateDetailTransaksi handles POST requests to create a new Detail Transaksi
func CreateDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	var detail models.DetailTransaksi
	err := json.NewDecoder(r.Body).Decode(&detail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO Detail_Transaksi (Id_Transaksi, Id_Buku, Jumlah, Harga) VALUES (?, ?, ?, ?)`
	_, err = config.DB.Exec(query, detail.Id_Transaksi, detail.Id_Buku, detail.Jumlah, detail.Harga)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(detail)
