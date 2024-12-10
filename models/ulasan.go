package models


type Ulasan struct {
	Id_Ulasan  int    `json:"id_ulasan"`
	Id_Pembeli int    `json:"id_pembeli"`
	Id_Buku    int    `json:"id_buku"`
	Rating     int    `json:"rating"`  // Rating buku, misalnya 1-5
	Ulasan     string `json:"ulasan"`  // Komentar atau ulasan pembeli
}
