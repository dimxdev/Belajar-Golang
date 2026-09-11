# 10 - Pointer

Contoh praktik: [`src/10-pointer.go`](../src/10-pointer.go)

Pointer adalah variabel yang isinya **alamat memori** dari variabel lain, bukan nilainya langsung.

## Dua Operator Utama

| Operator | Arti | Contoh |
|---|---|---|
| `&` | Ambil **alamat** dari variabel | `pNama := &nama` |
| `*` | Ambil/ubah **nilai** yang ditunjuk pointer (dereference) | `fmt.Println(*pNama)` |

## Bikin & Pakai Pointer

```go
nama := "Dimas"
pNama := &nama // & = "ambil alamat dari nama"

var umur int = 20
var pUmur *int = &umur // *int = "tipe pointer ke int"

fmt.Println(nama)   // Dimas       -> nilai aslinya
fmt.Println(pNama)  // 0xc0000...  -> alamat memori
fmt.Println(*pNama) // Dimas       -> * = "buka pointer, ambil nilainya"
```

## Mengubah Nilai Lewat Pointer

```go
alamat := "Jepara"
pAlamat := &alamat
*pAlamat = "Bandung" // ubah value asli lewat pointer

fmt.Println(alamat) // Bandung -> ikut berubah!
```
Karena `pAlamat` menunjuk ke alamat memori yang sama dengan `alamat`, mengubah `*pAlamat` = mengubah `alamat` aslinya juga.

## Pointer Kosong (`nil`)

```go
var role *int
fmt.Println(role)        // <nil>
fmt.Println(role == nil) // true
```
Pointer yang belum di-set menunjuk ke variabel manapun bernilai `nil` (zero value untuk pointer). Mengakses `*role` saat `role` masih `nil` akan menyebabkan **panic** (nil pointer dereference) — salah satu bug paling umum kalau lupa cek nil dulu.

## `new()` — Alokasi Pointer Kosong

```go
p := new(int)   // p adalah *int, menunjuk ke int dengan zero value (0)
*p = 5
```

## Kenapa Pointer Penting di Go

1. **Pass by reference ke function** — supaya function bisa mengubah variabel aslinya (lihat [12 - Pointer di Function](12-pointer-function.md)).
2. **Menghindari copy data besar** — struct besar yang dikirim sebagai pointer tidak perlu disalin seluruhnya.
3. **Pointer receiver di method** — supaya method bisa mengubah struct aslinya (lihat [13 - Pointer di Method](13-pointer-method.md)).
4. **Merepresentasikan "nilai opsional"** — misal field `*int` bisa membedakan antara "nilai 0" dan "belum diisi/nil", yang tidak bisa dibedakan kalau pakai `int` biasa.

## Poin Penting

- `&variabel` → dapatkan alamat (bikin pointer).
- `*pointer` → dapatkan/ubah nilai yang ditunjuk pointer.
- Selalu cek `pointer != nil` sebelum dereference kalau ada kemungkinan pointer itu belum di-set, untuk menghindari panic.
