# 📥 GET di `net/http`

Method `GET` dipakai buat **mengambil data**, tanpa mengirim body request. Data tambahan (kalau ada) biasanya dikirim lewat **query parameter** atau **path parameter**, bukan body.

## Contoh Paling Dasar

```go
type Profil struct {
	Nama string `json:"nama"`
	City string `json:"city"`
	Umur int    `json:"umur"`
}

func profilHandler(w http.ResponseWriter, r *http.Request) {
	data := Profil{
		Nama: "Dimas",
		City: "Jepara",
		Umur: 20,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
```

Alurnya: siapin data → set header `Content-Type` → `Encode` (convert struct jadi JSON, langsung ditulis ke response).

## Registrasi Route dengan Method Spesifik (Go 1.22+)

```go
http.HandleFunc("GET /profile", profilHandler)
```

Tanpa prefix method (`http.HandleFunc("/profile", profilHandler)`), route itu bakal nerima **semua** method (GET, POST, dll), bukan cuma GET. Prefix `"GET /profile"` bikin route ini cuma nyala buat method `GET`, request method lain ke path yang sama otomatis dapet `404`.

## Baca Query Parameter (`?key=value`)

Contoh: `GET /search?nama=budi&limit=10`

```go
func searchHandler(w http.ResponseWriter, r *http.Request) {
	nama := r.URL.Query().Get("nama")   // "budi"
	limit := r.URL.Query().Get("limit") // "10" (selalu string, perlu di-convert manual kalau butuh angka)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"nama":  nama,
		"limit": limit,
	})
}
```

`r.URL.Query().Get("key")` balikin string kosong `""` kalau key-nya nggak ada di URL — jadi kalau butuh validasi "wajib diisi", perlu dicek manual.

## Baca Path Parameter (Go 1.22+)

Contoh: `GET /user/5` (ambil user dengan ID 5)

```go
http.HandleFunc("GET /user/{id}", getUserHandler)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // "5", tipe string

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
```

`{id}` di pattern route itu jadi "slot" yang bisa diambil isinya lewat `r.PathValue("id")` — fitur ini baru ada sejak Go 1.22, sebelumnya harus pakai library routing pihak ketiga (atau parsing manual dari `r.URL.Path`) buat dapetin path parameter kayak gini.

## Response dengan Status Code Eksplisit

```go
func getProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	produk, ada := cariProduk(id)

	w.Header().Set("Content-Type", "application/json")
	if !ada {
		w.WriteHeader(http.StatusNotFound) // 404
		json.NewEncoder(w).Encode(map[string]string{"message": "produk tidak ditemukan"})
		return
	}

	w.WriteHeader(http.StatusOK) // 200, sebenarnya opsional karena ini default
	json.NewEncoder(w).Encode(produk)
}
```

Lihat daftar lengkap status code di [status-code.md](status-code.md).

## Poin Penting

- `GET` idealnya **tidak** mengubah data di server (idempotent) — kalau butuh membuat/mengubah/menghapus data, pakai method yang sesuai (`POST`/`PUT`/`DELETE`).
- `GET` **tidak punya body request** secara konvensi — data tambahan lewat query parameter (`?key=value`) atau path parameter (`/resource/{id}`), bukan `r.Body`.
- `r.URL.Query().Get(...)` dan `r.PathValue(...)` selalu balikin `string` — perlu `strconv.Atoi(...)` dkk kalau butuh tipe lain (misal `int`).
