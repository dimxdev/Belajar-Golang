# 📤 POST di `net/http`

Method `POST` dipakai buat **mengirim/membuat data baru**. Data yang dikirim client biasanya ada di **body request**, dalam format JSON, dan perlu di-*decode* dulu jadi struct Go sebelum bisa dipakai.

## Contoh Paling Dasar

```go
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
	w.WriteHeader(http.StatusCreated) // 201, default-nya 200 kalau gak ditulis manual
	json.NewEncoder(w).Encode(user)
}
```

Alurnya: `Decode` baca body JSON dari `r.Body` → isi ke struct `user` (lewat pointer `&user`) → kalau gagal, langsung `return` dengan status `400` → kalau berhasil, kirim balik konfirmasi.

## Registrasi Route dengan Method Spesifik (Go 1.22+)

```go
http.HandleFunc("POST /user", createUserHandler)
```

## Validasi Setelah Decode Berhasil

Decode berhasil cuma berarti **formatnya** JSON valid — bukan berarti **isinya** sesuai aturan bisnis. Validasi tambahan tetap perlu dicek manual, satu-satu, dengan response yang beda tiap kasus:

```go
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Status: "error", Message: "Data tidak valid"})
		return
	}

	if product.Harga <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		json.NewEncoder(w).Encode(Response{Status: "error", Message: "Harga harus lebih dari 0"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Status: "success", Data: product})
}
```

## Response Envelope (`status` + `data`)

Pola umum biar bentuk response konsisten di semua endpoint:

```go
type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
```

`Data any` (alias `interface{}`) dipakai karena isinya bisa beda-beda tergantung endpoint (`User`, `Product`, slice, dll). `omitempty` bikin field itu di-skip dari JSON kalau nilainya kosong/zero value.

## `Decode` vs `Encode` — Arah yang Berlawanan

| | Arah | Contoh |
|---|---|---|
| `Decode` | JSON masuk (dari `r.Body`) → jadi struct Go | `json.NewDecoder(r.Body).Decode(&user)` |
| `Encode` | struct Go → jadi JSON keluar (ke `w`) | `json.NewEncoder(w).Encode(user)` |

`Decode` butuh **pointer** (`&user`) karena dia perlu **mengisi** variabel yang dikasih, bukan cuma membaca/return nilai baru.

## Poin Penting

- Selalu cek `err != nil` setelah `Decode` — kalau diabaikan, request dengan body rusak/invalid tetap lanjut diproses dengan struct kosong (zero value), dan client bisa dapet response `200 OK` yang menyesatkan (seolah sukses padahal gagal).
- `w.Header().Set(...)` dan `w.WriteHeader(...)` harus dipanggil **sebelum** `Encode`/`Write` — begitu body mulai ditulis, header & status code otomatis ter-"lock".
- Status code yang umum dipakai di endpoint `POST`: `201` (berhasil buat data baru), `400` (body/format invalid), `422` (format valid tapi gagal validasi bisnis). Lihat daftar lengkap di [status-code.md](status-code.md).
