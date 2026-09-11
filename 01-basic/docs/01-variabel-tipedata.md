# 01 - Variabel & Tipe Data

Contoh praktik: [`src/01-variabel-tipedata.go`](../src/01-variabel-tipedata.go)

## Cara Deklarasi Variabel

**1. `var` + nama + tipe** — eksplisit, tipe wajib/boleh ditulis
```go
var name string = "Dimas"
var age int          // tanpa nilai awal -> otomatis zero value (0)
```

**2. Short declaration (`:=`)** — Go nebak tipe dari nilainya, cuma bisa dipakai di dalam function
```go
city := "Jepara"
```

**3. Deklarasi banyak variabel sekaligus**
```go
a, b := 1, 3
var x, y int = 1, 2
```

**4. Konstanta (`const`)** — nilai yang nggak boleh berubah setelah didefinisikan
```go
const Pi = 3.14
const MaxUser = 100
```

## Tipe Data Dasar

**Boolean**
| Tipe | Nilai |
|---|---|
| `bool` | `true` / `false` |

**Angka Bulat (Integer)**
| Tipe | Ukuran | Range |
|---|---|---|
| `int` | tergantung platform (32/64-bit) | paling umum dipakai |
| `int8`, `int16`, `int32`, `int64` | ukuran tetap sesuai nama | dipakai kalau butuh kontrol ukuran spesifik |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | unsigned (nggak bisa negatif) | dipakai kalau nilainya pasti selalu ≥ 0 |
| `byte` | alias dari `uint8` | biasa dipakai buat data biner/karakter |
| `rune` | alias dari `int32` | biasa dipakai buat merepresentasikan 1 karakter Unicode |

**Angka Desimal (Floating Point)**
| Tipe | Presisi |
|---|---|
| `float32` | presisi lebih rendah |
| `float64` | presisi lebih tinggi, **default** kalau nulis angka desimal tanpa tipe eksplisit |

**Teks**
| Tipe | Keterangan |
|---|---|
| `string` | teks, immutable (nggak bisa diubah sebagian, cuma bisa diganti seluruhnya) |

## Zero Value

Kalau variabel dideklarasi tanpa nilai awal, Go otomatis kasih **nilai default** sesuai tipenya:

| Tipe | Zero Value |
|---|---|
| `int`, `float64` | `0` |
| `string` | `""` (string kosong) |
| `bool` | `false` |
| pointer, slice, map, channel, func, interface | `nil` |

## Konversi Tipe (Type Conversion)

Go itu **strict**, nggak ada auto-convert antar tipe angka beda jenis:
```go
var i int = 10
var f float64 = float64(i) // wajib konversi eksplisit
```

## Poin Penting

- `var` cocok kalau nilai variabel mau di-set belakangan (bisa dideklarasi tanpa nilai awal).
- `:=` wajib langsung diisi nilai, dan cuma bisa dipakai di dalam function (bukan di level package/top-level).
- Go itu **statically typed** — sekali tipe ditentukan, variabel itu nggak bisa diisi tipe lain tanpa konversi eksplisit.
- Pilih tipe integer/float sesuai kebutuhan: default-nya `int` dan `float64` sudah cukup buat kebanyakan kasus, tipe lain (`int8`, `uint32`, dst) baru relevan kalau ada alasan spesifik (misal optimasi memory, kompatibilitas format data tertentu).
