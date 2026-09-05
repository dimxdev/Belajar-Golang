# ✅ Validasi di Gin (Tag `binding`)

Salah satu keunggulan Gin dibanding `net/http` polos: validasi data request bisa ditulis **langsung di tag struct**, bukan lewat `if` manual berulang-ulang. Ini jalan lewat library [`go-playground/validator`](https://github.com/go-playground/validator) yang udah terintegrasi otomatis di Gin.

```go
type Product struct {
	Nama string `json:"nama" binding:"required"`
}
```

Validasi ini **cuma jalan** kalau kamu bind pakai `c.ShouldBindJSON(&x)` / `c.Bind(&x)` — kalau decode manual pakai `encoding/json` biasa, tag `binding` diabaikan total.

---

## Daftar Tag Validasi

### Wajib Diisi

```go
Nama string `json:"nama" binding:"required"`
```
`required` — field wajib diisi, bukan zero value (`""`, `0`, `false`, slice/map kosong, dll dianggap "kosong").

### Angka: Batas Minimum/Maksimum

```go
Harga int `json:"harga" binding:"gt=0"`         // harus > 0
Stok  int `json:"stok" binding:"gte=0"`          // harus >= 0
Umur  int `json:"umur" binding:"gte=17,lte=100"` // antara 17-100
```

| Tag | Arti |
|---|---|
| `gt=n` | greater than — lebih besar dari n |
| `gte=n` | greater than or equal — lebih besar sama dengan n |
| `lt=n` | less than — lebih kecil dari n |
| `lte=n` | less than or equal — lebih kecil sama dengan n |

### String: Panjang Karakter

```go
Username string `json:"username" binding:"min=3,max=20"`
Kode     string `json:"kode" binding:"len=6"` // harus PERSIS 6 karakter
```

| Tag | Arti |
|---|---|
| `min=n` | panjang minimal n karakter (juga berlaku buat angka minimum & panjang slice) |
| `max=n` | panjang maksimal n karakter |
| `len=n` | panjang harus persis n karakter |

### Format Khusus

```go
Email   string `json:"email" binding:"required,email"`
Website string `json:"website" binding:"url"`
ID      string `json:"id" binding:"uuid"`
```

| Tag | Arti |
|---|---|
| `email` | harus format email valid |
| `url` | harus format URL valid |
| `uuid` | harus format UUID valid |
| `numeric` | isinya harus angka doang |
| `alpha` | isinya harus huruf doang |
| `alphanum` | isinya harus huruf + angka doang |

### Nilai Harus Salah Satu dari Pilihan

```go
Status string `json:"status" binding:"oneof=pending success failed"`
```
`oneof=...` — value-nya harus **persis salah satu** dari daftar yang dikasih (dipisah spasi), selain itu ditolak.

### Kombinasi Banyak Aturan (Dipisah Koma)

```go
type Product struct {
	Nama  string  `json:"nama" binding:"required,min=3,max=50"`
	Harga float64 `json:"harga" binding:"required,gt=0"`
	Email string  `json:"email" binding:"required,email"`
}
```

### Field Opsional (Boleh Kosong)

```go
Motto string `json:"motto,omitempty"` // gak ada tag binding sama sekali -> gak divalidasi
```
Kalau field nggak dikasih tag `binding`, artinya bebas — boleh diisi apa aja, termasuk kosong.

---

## Contoh Lengkap: Pakai Validasi Default

```go
type Product struct {
	Nama  string  `json:"nama" binding:"required,min=3"`
	Harga float64 `json:"harga" binding:"required,gt=0"`
	Stok  int     `json:"stok" binding:"gte=0"`
}

func createProductHandler(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success", "data": product})
}
```

Kalau client kirim `{"nama": "AB", "harga": -5}` (nama cuma 2 karakter, harga negatif), `ShouldBindJSON` langsung gagal duluan — kamu **nggak perlu** nulis `if product.Harga <= 0 { ... }` manual sama sekali.

Tapi `err.Error()` di atas ngeluarin pesan default yang kurang enak dibaca, contoh:
```
Key: 'Product.Nama' Error:Field validation for 'Nama' failed on the 'min' tag
Key: 'Product.Harga' Error:Field validation for 'Harga' failed on the 'gt' tag
```
Teknis banget, bahasa Inggris, dan bocorin nama field/tag internal — kurang cocok dikirim langsung ke user aplikasi.

---

## Custom Pesan Error

Buat bikin pesan lebih ramah, "bongkar" error jadi tipe `validator.ValidationErrors`, terus mapping manual per-field:

```go
import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func createProductHandler(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorMessages := make(map[string]string)
			for _, fe := range ve {
				errorMessages[fe.Field()] = pesanCustom(fe)
			}
			c.JSON(400, gin.H{"status": "error", "errors": errorMessages})
			return
		}
		c.JSON(400, gin.H{"status": "error", "message": "Data tidak valid"})
		return
	}

	c.JSON(201, gin.H{"status": "success", "data": product})
}

func pesanCustom(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " wajib diisi"
	case "gt":
		return fe.Field() + " harus lebih dari " + fe.Param()
	case "gte":
		return fe.Field() + " harus lebih dari atau sama dengan " + fe.Param()
	case "min":
		return fe.Field() + " minimal " + fe.Param() + " karakter"
	case "max":
		return fe.Field() + " maksimal " + fe.Param() + " karakter"
	case "email":
		return fe.Field() + " harus berupa email yang valid"
	case "oneof":
		return fe.Field() + " harus salah satu dari: " + fe.Param()
	default:
		return fe.Field() + " tidak valid"
	}
}
```

Hasilnya jadi lebih ramah dibaca:
```json
{
	"status": "error",
	"errors": {
		"Nama": "Nama minimal 3 karakter",
		"Harga": "Harga harus lebih dari 0"
	}
}
```

### Penjelasan Method di `validator.FieldError`

| Method | Fungsi | Contoh Hasil |
|---|---|---|
| `fe.Field()` | nama field yang gagal validasi | `"Harga"` |
| `fe.Tag()` | tag validasi mana yang gagal | `"gt"` |
| `fe.Param()` | nilai parameter dari tag itu | `"0"` (dari `gt=0`) |
| `fe.Value()` | nilai yang dikirim client (yang gagal validasi) | `-5` |

---

## Kapan Perlu Custom Pesan, Kapan Nggak?

| Situasi | Perlu custom? |
|---|---|
| Belajar / project kecil / debugging | Nggak perlu, `err.Error()` polos udah cukup |
| API yang dipakai frontend/user beneran | Disarankan custom, biar pesan lebih jelas & nggak bocorin detail teknis internal |
| Tim besar / produk publik | Custom + mungkin bikin helper reusable (function `pesanCustom` dipakai ulang di semua handler) |

---

## Poin Penting

- Tag `binding` cuma jalan lewat `ShouldBindJSON`/`ShouldBind`/`Bind` — bukan lewat `encoding/json` biasa.
- Validasi generik (wajib isi, format, range angka) bisa dipindah ke tag struct; validasi **spesifik ke logika bisnis** (misal "stok tidak boleh melebihi kapasitas gudang") tetap harus dicek manual di handler.
- Pesan error default berguna buat development, tapi sebaiknya di-custom kalau API-nya bakal dipakai user/frontend beneran.
- Referensi terkait: [POST.md](POST.md), [gin-context.md](gin-context.md).
