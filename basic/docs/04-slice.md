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

## Poin Penting

- `append` bisa mengubah slice asli **atau** bikin array baru di belakang layar (tergantung kapasitas) — selalu tampung hasilnya ke variabel.
- Slicing `[a:b]`: index awal termasuk, index akhir **tidak** termasuk.
- Slice itu sebenarnya "view" ke array di belakang layar (punya pointer, length, capacity) — dua slice bisa berbagi array yang sama, jadi hati-hati efek samping kalau slice-nya dimodifikasi dari 2 tempat berbeda.
