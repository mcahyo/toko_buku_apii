package main

import (
	"fmt"
	"log"
	"net/http"
	"toko_buku_api/config"
	"toko_buku_api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	config.ConnectDB()

	// Setup the router
	r := mux.NewRouter()

	// Pembeli routes
	r.HandleFunc("/api/pembeli", controllers.GetPembeli).Methods("GET")
	r.HandleFunc("/api/pembeli/{id}", controllers.GetPembeliByID).Methods("GET")
	r.HandleFunc("/api/pembeli", controllers.CreatePembeli).Methods("POST")

	// Buku routes
	r.HandleFunc("/api/buku", controllers.GetBuku).Methods("GET")
	r.HandleFunc("/api/buku/{id}", controllers.GetBukuByID).Methods("GET")
	r.HandleFunc("/api/buku", controllers.CreateBuku).Methods("POST")

	// Ulasan routes
	r.HandleFunc("/api/ulasan", controllers.CreateUlasan).Methods("POST")

	// Transaksi routes
	r.HandleFunc("/api/transaksi", controllers.CreateTransaksi).Methods("POST")

	// Detail Transaksi routes
	r.HandleFunc("/api/detail-transaksi", controllers.CreateDetailTransaksi).Methods("POST")

	// Start the server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
