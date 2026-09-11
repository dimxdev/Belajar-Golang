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
	w.WriteHeader(http.StatusCreated) //ngasih status code klo gk ditulis manualnya 200
	json.NewEncoder(w).Encode(user)
}


type Product struct {
	Name string `json:"name"`
	Harga float64 `json:"harga"`
	Stock int `json:"stock"`
}

type Response struct {
	Status string `json:"status"`
	Data any `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status: "error",
			Message: "Data yg dimasukkan tidak valid",
		})

		return
	}

	if product.Harga < 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{
			Status: "error",
			Message: "Harga yg dimasukkan tidak boleh kurang dari 0",
		})

		return
	} 

	if product.Stock < 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{
			Status: "error",
			Message: "stock yg dimasukkan tidak boleh kurang dari 0",
		})

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{
		Status: "succes",
		Data: product,
	})
}


func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/profile", profilHandler)
	http.HandleFunc("/user", createUserHandler)
	http.HandleFunc("POST /product", createProductHandler)

	fmt.Println("Server jalan di http://localhost:1501")
	http.ListenAndServe(":1501", nil)
}
