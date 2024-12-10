package models

// DetailTransaksi represents the structure of the Detail_Transaksi table
type DetailTransaksi struct {
	Id_Detail_Transaksi int    `json:"id_detail_transaksi"`
	Id_Transaksi        int    `json:"id_transaksi"`
	Id_Buku             int    `json:"id_buku"`
	Jumlah             int    `json:"jumlah"`
	Harga              float64 `json:"harga"` // Harga per buku pada saat transaksi
}