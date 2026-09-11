# 04 - Slice

Contoh praktik: [`src/04-slice.go`](../src/04-slice.go)

## Slice vs Array

Go punya 2 tipe koleksi data berurutan: **array** (ukuran tetap/fix) dan **slice** (ukuran dinamis, lebih fleksibel & jauh lebih sering dipakai).

```go
var arr [3]string        // array, ukuran fix 3, gak bisa nambah/kurang
siswa := []string{"A", "B", "C"} // slice, ukuran dinamis
```

## Deklarasi & Akses

```go
siswa := []string{"Dimas", "Julian", "Rifqi"}
fmt.Println(siswa[0])   // akses per-index (mulai dari 0)
fmt.Println(len(siswa)) // jumlah elemen
```

Bikin slice kosong dengan `make`:
```go
angka := make([]int, 0)      // slice kosong
angka2 := make([]int, 5)     // slice isi 5 elemen, semua zero value
```

## Operasi Umum

**Nambah elemen (`append`)** — selalu return slice baru, wajib ditampung ulang:
```go
siswa = append(siswa, "Safril")
```

**Slicing (ambil sebagian)** — index awal termasuk, index akhir **tidak termasuk**:
```go
potongan := siswa[1:3] // ambil index 1 sampai sebelum index 3
```

**Gabung 2 slice** — `...` "membongkar" elemen slice satu-satu:
```go
gabungan := append(siswa, siswaB...)
```

**Hapus elemen** (Go nggak punya function bawaan, biasa dikombinasi slicing + append):
```go
i := 2
siswa = append(siswa[:i], siswa[i+1:]...) // hapus elemen index ke-2
```

**Copy slice**
```go
tujuan := make([]int, len(sumber))
copy(tujuan, sumber)
```

## Loop Slice (`range`)

```go
for i, fruit := range fruits {
	fmt.Printf("%d. %s\n", i, fruit) // index + value
}
for _, fruit := range fruits {
	fmt.Println(fruit)               // value doang, index diabaikan
}
for i := range fruits {
	fmt.Println(i)                   // index doang
}
```

## Slice Tipe Campuran & Slice of Struct

```go
campur := []any{"Budi", 123, true, 3.14} // any (interface{}) buat tipe campuran

type Barang struct {
	Name  string
	Harga int
}
barangs := []Barang{
	{Name: "Keyboard", Harga: 100000},
	{Name: "Laptop", Harga: 1000000},
}
```
Pola slice-of-struct ini yang paling sering dipakai buat representasi "list data" — hasil query database, atau response API berbentuk array of object.

## Slice sebagai Tipe Field di Struct

Slice juga bisa jadi **tipe field** di dalam struct — dipakai kalau satu field itu perlu nampung **banyak nilai sekaligus**, bukan cuma 1 nilai tunggal.

```go
type Config struct {
	AllowOrigins []string // field ini nampung banyak domain sekaligus
	AllowMethods []string // field ini nampung banyak method sekaligus
	AllowHeaders []string
}
```

Cara isinya, gabungan struct literal + slice literal:
```go
config := Config{
	AllowOrigins: []string{"http://localhost:3000"},                 // 1 elemen
	AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},           // 4 elemen
	AllowHeaders: []string{"Content-Type", "Authorization"},
}
```

Kenapa dipilih `[]string` (slice), bukan `string` tunggal? Karena secara logika, nilainya **bisa lebih dari 1** — misal domain yang diizinkan bisa banyak, method HTTP yang diizinkan juga bisa banyak. Kalau dipaksa pakai `string` tunggal, field itu cuma bisa nampung 1 nilai doang, nggak fleksibel.

Field struct juga bisa diisi slice of struct, kalau butuh nampung banyak data yang lebih kompleks dari sekadar string/int:
```go
type Toko struct {
	Nama    string
	Produk  []Barang // slice of struct, lihat definisi Barang di atas
}

toko := Toko{
	Nama: "Toko Budi",
	Produk: []Barang{
		{Name: "Keyboard", Harga: 100000},
		{Name: "Laptop", Harga: 1000000},
	},
}
```

Ini juga berlaku buat tipe data lain, bukan cuma slice — field struct bisa diisi tipe **apapun** yang valid di Go: `map`, struct lain (nested), pointer, bahkan `func` atau `chan`:
```go
type Server struct {
	Nama    string            // tipe dasar
	Domains []string          // slice
	Headers map[string]string // map
	Alamat  Alamat            // struct lain (nested)
	Owner   *User             // pointer
}
```
Aturan generalnya: tipe field struct itu bebas, tinggal disesuaikan sama bentuk data yang mau ditampung — bukan cuma terbatas ke `string`/`int`/`bool`.

## Poin Penting

- `append` bisa mengubah slice asli **atau** bikin array baru di belakang layar (tergantung kapasitas) — selalu tampung hasilnya ke variabel.
- Slicing `[a:b]`: index awal termasuk, index akhir **tidak** termasuk.
- Slice itu sebenarnya "view" ke array di belakang layar (punya pointer, length, capacity) — dua slice bisa berbagi array yang sama, jadi hati-hati efek samping kalau slice-nya dimodifikasi dari 2 tempat berbeda.
