# 🎯 `gin.Context` (si `c`)

`*gin.Context` adalah parameter tunggal yang masuk ke tiap handler Gin — satu objek yang gabungin semua kebutuhan **baca request** dan **tulis response**, gantiin `w http.ResponseWriter` + `r *http.Request` yang terpisah di `net/http`.

```go
func profilHandler(c *gin.Context) {
	// c dipakai buat baca request DAN nulis response
}
```

## Nulis Response

| Method | Fungsi |
|---|---|
| `c.JSON(code, data)` | kirim response JSON (paling sering dipakai) |
| `c.String(code, format, args...)` | kirim response teks biasa |
| `c.XML(code, data)` | kirim response XML |
| `c.HTML(code, template, data)` | render & kirim HTML template |
| `c.Data(code, contentType, bytes)` | kirim response raw bytes, bebas format |
| `c.File(path)` | kirim isi file (download/serve gambar, dll) |
| `c.Redirect(code, url)` | redirect ke URL lain |
| `c.Status(code)` | set status code doang, tanpa body |

## Baca Data dari Request

| Method | Fungsi |
|---|---|
| `c.Param("id")` | ambil path parameter, misal dari route `/user/:id` |
| `c.Query("nama")` | ambil query parameter, misal dari `?nama=budi` |
| `c.DefaultQuery("limit", "10")` | ambil query parameter, dengan nilai default kalau kosong |
| `c.PostForm("nama")` | ambil data dari form (`application/x-www-form-urlencoded`) |
| `c.ShouldBindJSON(&struct)` | decode body JSON ke struct, return `error` kalau gagal |
| `c.Bind(&struct)` | mirip `ShouldBindJSON`, tapi otomatis kirim response `400` sendiri kalau gagal (kurang fleksibel) |
| `c.GetHeader("Authorization")` | ambil value dari header request |
| `c.Request` | akses objek `*http.Request` mentah, kalau butuh sesuatu di luar shortcut Gin |

## Utility Lain

| Method | Fungsi |
|---|---|
| `c.Set("key", value)` / `c.Get("key")` | simpan/ambil data sementara, dipakai buat oper data antar middleware & handler dalam 1 siklus request |
| `c.Next()` | lanjutkan ke middleware/handler berikutnya (dipakai di dalam middleware) |
| `c.Abort()` | hentikan proses, jangan lanjut ke handler berikutnya (misal auth gagal) |
| `c.FullPath()` | ambil pattern route yang match, misal `/user/:id` |

## `gin.H` — Shortcut Bikin Map buat JSON

```go
c.JSON(200, gin.H{
	"status": "success",
	"data":   user,
})
```
`gin.H` itu cuma alias dari `map[string]any` (alias dari `map[string]interface{}`) — dibikin biar nulisnya lebih pendek daripada `map[string]interface{}{...}` berulang-ulang. Fungsinya sama persis kayak bikin `Response` struct yang dibahas di [POST.md net/docs](../../net/docs/POST.md), tapi lebih cepat ditulis untuk kasus yang nggak butuh struct formal.

## Contoh Pola Umum: Baca Body + Validasi + Response

```go
func createUserHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": "Data tidak valid"})
		return
	}

	if user.Umur <= 0 {
		c.JSON(422, gin.H{"status": "error", "message": "Umur harus lebih dari 0"})
		return
	}

	c.JSON(201, gin.H{"status": "success", "data": user})
}
```
Pola ini identik dengan pola `Decode` + validasi berlapis yang dibahas di [POST.md net/docs](../../net/docs/POST.md), cuma versi Gin lebih ringkes.

## Poin Penting

- `c.JSON(code, data)` itu melakukan **2 hal sekaligus**: set status code + tulis body response — biasanya jadi baris terakhir sebelum handler selesai.
- `c.ShouldBindJSON` vs `c.Bind`: pakai `ShouldBindJSON` kalau mau kontrol sendiri response error-nya (lebih fleksibel, lebih umum dipakai).
- Semua daftar status code (`200`, `201`, `400`, dst) yang dipakai di `c.JSON(code, ...)` sama persis dengan yang dibahas di [status-code.md net/docs](../../net/docs/status-code.md) — Gin tidak punya sistem status code sendiri, tetap pakai standar HTTP yang sama.
