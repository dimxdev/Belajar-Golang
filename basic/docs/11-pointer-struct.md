# 11 - Pointer di Struct

Contoh praktik: [`src/11-pointer-struct.go`](../src/11-pointer-struct.go)

## Pointer ke Struct & Akses Field

```go
type Kucing struct {
	Nama  string
	umur  int
}

kucing1 := Kucing{Nama: "Alex", umur: 3}
pKucing1 := &kucing1   // pointer ke kucing1
pKucing1.Nama = "Boni" // ubah field langsung lewat pointer
```

Akses field lewat pointer struct ada 2 versi yang setara:
```go
fmt.Println((*pKucing1).umur) // versi eksplisit: dereference dulu, baru akses field
fmt.Println(pKucing1.umur)    // versi singkat: Go otomatis dereference-in
```
Go otomatis menerjemahkan `pKucing1.umur` menjadi `(*pKucing1).umur` di belakang layar — makanya versi singkat yang umum dipakai.

## Struct Nested + Pointer

```go
type Alamat struct {
	Kota string
}
type Pekerja struct {
	Nama   string
	Alamat Alamat
}

pekerja1 := Pekerja{Nama: "Budi", Alamat: Alamat{Kota: "Jepara"}}
pPekerja1 := &pekerja1
pPekerja1.Alamat.Kota = "Semarang" // akses field nested lewat pointer tetap mudah
```

## Struct dengan Field Bertipe Pointer

```go
type Anjing struct {
	Name string
	age  *int // pointer, bukan int langsung
}

umur := 30
anjing1 := Anjing{Name: "Lutsi", age: &umur}
```
Berguna kalau butuh membedakan "nilai 0" dari "belum diisi" (`nil`), atau field itu perlu "terhubung" ke variabel di luar struct.

## Slice/Array Berisi Pointer ke Struct

Pola umum lain: koleksi berisi pointer ke struct, bukan struct langsung — supaya perubahan pada satu elemen tercermin di mana pun elemen itu direferensikan:
```go
kucings := []*Kucing{
	{Nama: "Alex"},
	{Nama: "Boni"},
}
kucings[0].Nama = "Cimol" // langsung ubah, tanpa perlu ambil ulang index
```

## Poin Penting

- Akses field lewat pointer struct **tidak perlu** dereference manual (`(*p).field`), cukup `p.field` — Go otomatis mengonversi.
- Pola `pStruct := &struct1` sangat umum dipakai saat struct perlu diubah lewat function atau method (lanjut ke [12](12-pointer-function.md) & [13](13-pointer-method.md)).
- Berbagi struct lewat pointer antar bagian kode berarti semua pihak melihat perubahan yang sama — perhatikan efek sampingnya kalau tidak diinginkan (kadang perlu bikin salinan eksplisit).
