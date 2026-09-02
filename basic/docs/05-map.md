# 05 - Map

Contoh praktik: [`src/05-map.go`](../src/05-map.go)

Map adalah struktur data **key-value**, mirip object di JS atau dictionary di Python — cocok buat lookup cepat berdasarkan key.

## Deklarasi & Akses

```go
nilai := map[string]int{
	"Dimas": 100,
	"Azis":  98,
}
fmt.Println(nilai["Dimas"]) // akses value lewat key
```

Bikin map kosong dengan `make`:
```go
data := make(map[string]int)
```

## Tambah / Update / Hapus

```go
nilai["Nabila"] = 99  // key baru -> nambah
nilai["Dimas"] = 97   // key udah ada -> update/timpa
delete(nilai, "Azis") // hapus key
```

## Cek Apakah Key Ada ("comma ok" idiom)

```go
value, exists := nilai["Andi"]
if !exists {
	fmt.Println("Andi tidak terdaftar")
}
```
Penting: akses key yang **nggak ada** tanpa cek `exists` tidak error, tapi balikin **zero value** (`0` buat map `int`) — bisa nyasar dianggap valid padahal key-nya emang nggak ada.

## Loop Map

```go
for key, value := range nilai {
	fmt.Printf("%s: %d\n", key, value)
}
```

## Map dengan Value Berupa Struct / Slice

```go
type User struct {
	Nama string
	Umur int
}

users := map[string]User{
	"u1": {Nama: "Dimas", Umur: 20},
}
```

## `len()` pada Map

```go
fmt.Println(len(nilai)) // jumlah key-value pair
```

## Poin Penting

- **Urutan map nggak dijamin konsisten** kalau di-loop — beda run bisa beda urutan (ini disengaja oleh Go). Kalau butuh urutan tertentu, sort key-nya dulu secara manual.
- Map key harus bertipe **comparable** (string, int, dll) — nggak bisa pakai slice atau map lain sebagai key.
- Map cocok buat lookup cepat berdasarkan identitas unik (misal cek "apakah user dengan ID ini ada?"), sedangkan slice lebih cocok buat data berurutan/list.
