# 12 - Pointer di Function

Contoh praktik: [`src/12-pointer-function.go`](../src/12-pointer-function.go)

## Kenapa Butuh Pointer di Parameter Function?

Secara default, Go itu **pass by value** — mengirim variabel ke function berarti mengirim **salinannya**, bukan variabel aslinya. Perubahan di dalam function **tidak** memengaruhi variabel asli di luar, kecuali dikirim dalam bentuk **pointer**.

## Contoh: Tipe Dasar

```go
func ubahNilai(nilaiAwal *int) {
	*nilaiAwal = 100
}

nilai := 10
ubahNilai(&nilai) // kirim alamat, bukan salinan
fmt.Println(nilai) // 100, ikut berubah!
```
Kalau parameternya `int` biasa (bukan `*int`), perubahan di dalam function cuma mengubah **salinan lokal**, variabel aslinya tetap `10`.

## Contoh: Struct

```go
type Food struct {
	nama string
	rate int
}

func ubahNamaFood(f *Food) {
	f.rate = 10 // Go otomatis dereference, sama efeknya dengan (*f).rate = 10
}

food1 := Food{nama: "Gule", rate: 9}
ubahNamaFood(&food1)
fmt.Println(food1) // rate berubah jadi 10
```

## Kapan Struct Dikirim sebagai Value vs Pointer?

| Kirim sebagai | Efek | Kapan dipakai |
|---|---|---|
| `Tipe` (value) | function dapat salinan, perubahan tidak ngaruh ke asli | struct kecil, data cuma dibaca |
| `*Tipe` (pointer) | function dapat alamat, perubahan ngaruh ke asli | struct besar (hindari copy), atau memang butuh mengubah struct asli |

## Slice & Map: Pengecualian

Slice dan map **tidak perlu** dikirim sebagai pointer untuk bisa "diubah" isinya, karena internal-nya sudah membawa referensi ke data:
```go
func tambahElemen(s []int) {
	s[0] = 999 // elemen yang sudah ada bisa diubah tanpa pointer
}
```
Tapi kalau butuh **mengubah panjang/reassign slice itu sendiri** (misal `append` di dalam function agar terlihat di luar), tetap butuh pointer (`*[]int`) atau `return` slice barunya.

## Poin Penting

- **Pass by value** (default Go) = function menerima **salinan**, perubahan di dalam tidak berdampak ke luar.
- **Pass by pointer** (`*Tipe`, dikirim dengan `&variabel`) = function menerima **alamat**, perubahan di dalam benar-benar mengubah variabel aslinya.
- Prinsip yang sama persis dipakai di pointer receiver pada method (lihat [13 - Pointer di Method](13-pointer-method.md)).
