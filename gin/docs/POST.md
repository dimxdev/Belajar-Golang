# 📤 POST di Gin

Sama seperti [POST.md net/docs](../../net/docs/POST.md): method `POST` dipakai buat **mengirim/membuat data baru**, datanya ada di body request (biasanya JSON), perlu di-*bind* dulu jadi struct Go.

## Contoh Paling Dasar

```go
type User struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

func createUserHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": "Data tidak valid"})
		return
	}

	c.JSON(201, gin.H{"status": "success", "data": user})
}
```

Alurnya: `ShouldBindJSON` baca body JSON dari request → isi ke struct `user` (lewat pointer `&user`) → kalau gagal, `return` dengan status `400` → kalau berhasil, kirim balik konfirmasi.

## Registrasi Route

```go
router.POST("/user", createUserHandler)
```

## Validasi Setelah Bind Berhasil

Bind berhasil cuma berarti **formatnya** JSON valid — bukan berarti isinya sesuai aturan bisnis. Validasi tambahan tetap perlu dicek manual:

```go
func createProductHandler(c *gin.Context) {
	var product Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": "Data tidak valid"})
		return
	}

	if product.Harga <= 0 {
		c.JSON(422, gin.H{"status": "error", "message": "Harga harus lebih dari 0"})
		return
	}

	c.JSON(201, gin.H{"status": "success", "data": product})
}
```

## Validasi Otomatis Lewat Tag `binding`

Ini yang bikin Gin lebih unggul dibanding `net/http` polos — banyak validasi bisa langsung ditulis di tag struct, tanpa `if` manual:

```go
type Product struct {
	Nama  string  `json:"nama" binding:"required"`
	Harga float64 `json:"harga" binding:"required,gt=0"`
}
```
- `required` — field wajib diisi (bukan zero value).
- `gt=0` — nilainya harus lebih besar dari 0.

Kalau validasi ini gagal, `c.ShouldBindJSON(&product)` langsung return `error` duluan, sebelum sempat masuk ke pengecekan manual — jadi validasi umum (wajib isi, minimal/maksimal, format email, dll) bisa dipindah ke tag struct, dan `if` manual cukup dipakai buat aturan bisnis yang lebih spesifik.

## `Decode`/`Encode` (net/http) vs `Bind`/`c.JSON` (Gin)

| | `net/http` | Gin |
|---|---|---|
| JSON masuk → struct | `json.NewDecoder(r.Body).Decode(&user)` | `c.ShouldBindJSON(&user)` |
| struct → JSON keluar | `json.NewEncoder(w).Encode(user)` | `c.JSON(code, user)` |

Konsepnya identik, `ShouldBindJSON` cuma nambah bonus validasi lewat tag `binding` yang nggak ada di `encoding/json` biasa.

## Poin Penting

- Selalu cek `err != nil` setelah `ShouldBindJSON` — sama pentingnya dengan `Decode` di `net/http`, biar request dengan body invalid nggak lanjut diproses.
- Tag `binding` di struct adalah keunggulan Gin buat validasi dasar (required, format, range angka) tanpa nulis `if` manual berulang-ulang.
- Status code yang umum di endpoint `POST`: `201` (berhasil buat data baru), `400` (body/format invalid), `422` (format valid tapi gagal validasi bisnis) — lihat [status-code.md net/docs](../../net/docs/status-code.md).
