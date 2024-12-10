package models


type Pembeli struct {
	Id_Pembeli   int    `json:"id_pembeli"`
	Nama_Pembeli string `json:"nama_pembeli"`
	Kata_Sandi   string `json:"kata_sandi"`
	Alamat_Email string `json:"alamat_email"`
}