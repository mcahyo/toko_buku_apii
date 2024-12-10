package models

import "time"


type Transaksi struct {
	Id_Transaksi  int       `json:"id_transaksi"`
	Id_Pembeli    int       `json:"id_pembeli"`
	Total_Harga   float64   `json:"total_harga"`
	Tanggal_Transaksi time.Time `json:"tanggal_transaksi"`
	Status        string    `json:"status"` // Contoh: "pending", "completed"
}
