package models

import "time"


type Buku struct {
	Id_Buku      int       `json:"id_buku"`
	Judul        string    `json:"judul"`
	Penulis      string    `json:"penulis"`
	Jenis        string    `json:"jenis"`
	ISBN         string    `json:"isbn"`
	Tahun_Terbit time.Time `json:"tahun_terbit"`
	Harga        float64   `json:"harga"`
	Stok         int       `json:"stok"`
}
