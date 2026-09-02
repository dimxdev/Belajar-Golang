# 03 - Struct

Contoh praktik: [`src/03-struct.go`](../src/03-struct.go)

Struct adalah cara bikin **tipe data custom** yang ngumpulin beberapa field (dengan tipe apapun, boleh beda-beda) jadi satu kesatuan. Konsepnya mirip "object" di bahasa lain, tapi tanpa class/inheritance.

## Definisi & Inisialisasi

```go
type User struct {
	Name  string
	Email string
	Age   int
}
```

**Cara isi berurutan** (harus sesuai urutan field, rawan salah kalau field-nya banyak):
```go
user1 := User{"Dimas", "dimas@mail.com", 20}
```

**Cara isi pakai nama field** (lebih jelas, urutan bebas — disarankan):
```go
user2 := User{Name: "Acep", Email: "acep@mail.com", Age: 25}
```

**Tanpa inisialisasi** — field otomatis keisi zero value:
```go
var user3 User // {"" "" 0}
```

## Exported vs Unexported Field

```go
type Product struct {
	Name  string // exported (huruf besar) -> bisa diakses dari package lain, ikut di-encode JSON
	stock int    // unexported (huruf kecil) -> cuma bisa diakses dari package ini sendiri
}
```
Field unexported juga **tidak** ikut ke-encode kalau struct-nya di-convert ke JSON (`encoding/json`).

## Struct Nested (Struct di Dalam Struct)

```go
type Alamat struct {
	Kota string
}

type Pegawai struct {
	Nama   string
	Alamat Alamat // struct di dalam struct
}

p := Pegawai{Nama: "Budi", Alamat: Alamat{Kota: "Jepara"}}
fmt.Println(p.Alamat.Kota)
```

## Anonymous Struct

Struct yang dipakai sekali aja tanpa perlu didefinisikan tipe barunya:
```go
point := struct {
	X, Y int
}{X: 1, Y: 2}
```

## Struct Embedding (Semacam "Pewarisan" ala Go)

```go
type Hewan struct {
	Nama string
}

type Kucing struct {
	Hewan // embedded struct, tanpa nama field
	Warna string
}

k := Kucing{Hewan: Hewan{Nama: "Alex"}, Warna: "Oren"}
fmt.Println(k.Nama) // bisa akses langsung field dari Hewan
```

## Tag Struct (Metadata Field)

Dipakai buat ngasih instruksi ke package lain (paling umum: `encoding/json`) soal gimana field itu diproses:
```go
type Profil struct {
	Nama string `json:"nama"`
}
```

## Perbandingan Struct

Struct bisa dibandingkan langsung pakai `==` **kalau** semua field-nya juga comparable (bukan slice/map/func):
```go
user1 == user2 // valid kalau semua field User comparable
```

## Poin Penting

- Struct adalah fondasi buat method (lihat 09), pointer ke struct (11), dan request/response body di REST API.
- Field diawali huruf besar = exported/public, huruf kecil = unexported/private ke package.
- Go tidak punya class/inheritance seperti OOP klasik — "reuse" struktur dilakukan lewat embedding, bukan pewarisan.
