package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Dasar HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Halo, ini server Go!")
}

// GET
type Profil struct {
	Nama string `json:"nama"`
	City string `json:"city"`
	Umur int `json:"umur"`
	Hobi string //auto pake nama asli waktu diubah jadi json
	Motto string `json:"motto,omitempty"` //kalo kosong dia gak ikut diencode ke json
}

func profilHandler(w http.ResponseWriter, r *http.Request) {
	data := Profil{
		Nama: "Dimas Brotowali H",
		City: "Jepara",
		Umur: 20,
		Hobi: "Ngoding",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// POST
type User struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/profile", profilHandler)
	http.HandleFunc("/user", createUserHandler)

	fmt.Println("Server jalan di http://localhost:1501")
	http.ListenAndServe(":1501", nil)
}
