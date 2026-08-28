package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Halo, ini server Go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server jalan di http://localhost:1501")
	http.ListenAndServe(":1501", nil)
}
