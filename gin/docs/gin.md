# 🍸 Gin Itu Sebenernya Apa Sih?

## TL;DR

Gin adalah **framework HTTP** buat Go — bukan bagian dari standard library (jadi butuh `go get github.com/gin-gonic/gin` dulu). Gin **dibangun di atas** [`net/http`](../../net/docs/net-http.md), bukan pengganti totalnya — tugasnya "membungkus" hal-hal yang di `net/http` butuh ditulis manual berulang-ulang (set header, encode JSON, routing per-method, dll) jadi method-method siap pakai yang lebih ringkes.

## Kenapa Orang Pilih Gin Dibanding `net/http` Polos?

1. **Routing lebih ekspresif** — method HTTP (`GET`, `POST`, dll) langsung jadi nama function router (`router.GET(...)`), plus path parameter (`router.GET("/user/:id", ...)`) yang gampang diambil pakai `c.Param("id")`.
2. **`gin.Context` (`c`)** — satu objek yang gabungin semua kebutuhan baca request + tulis response, gantiin `w http.ResponseWriter` + `r *http.Request` yang terpisah di `net/http`.
3. **Auto binding JSON** — `c.ShouldBindJSON(&struct)` gantiin `json.NewDecoder(r.Body).Decode(&struct)`, satu baris, sekalian validasi format.
4. **Middleware ecosystem** — banyak middleware siap pakai (logging, auth, CORS, rate limiting) tinggal `router.Use(...)`.
5. **Performa** — Gin dikenal salah satu framework Go tercepat, karena routing-nya pakai struktur data trie (radix tree) yang efisien.

## Perbandingan Cepat

| | `net/http` | Gin |
|---|---|---|
| Bagian dari standard library? | Ya | Tidak, perlu `go get` |
| Set response JSON | `w.Header().Set(...)` + `json.NewEncoder(w).Encode(...)` | `c.JSON(code, data)` |
| Baca body JSON | `json.NewDecoder(r.Body).Decode(&x)` | `c.ShouldBindJSON(&x)` |
| Routing per-method | `"GET /path"` (Go 1.22+) | `router.GET("/path", handler)` |
| Path parameter | `r.PathValue("id")` (Go 1.22+) | `c.Param("id")` |
| Middleware | manual/`http.Handler` wrapping | `router.Use(...)`, ekosistem siap pakai |

## Struktur Dasar Project Gin

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default() // bikin router, sekalian pasang middleware default (logger + recovery)

	router.GET("/profile", profilHandler)

	router.Run(":8000") // nyalain server, setara http.ListenAndServe
}
```

`gin.Default()` vs `gin.New()`: `Default()` otomatis pasang 2 middleware bawaan (logger request ke terminal + recovery dari panic biar server nggak crash total), sedangkan `New()` bikin router polos tanpa middleware apapun.

## Poin Penting

- Gin **tidak** menggantikan konsep-konsep dasar HTTP yang udah kamu pelajari di `net/http` — semua tetap sama (method, status code, header, JSON encode/decode). Gin cuma nyingkat cara nulisnya.
- Karena itu, urutan belajar `net/http` dulu baru Gin itu **tepat** — kamu udah paham "apa yang sebenarnya terjadi" di balik tiap shortcut Gin.
- Referensi lanjut: [GET.md](GET.md), [POST.md](POST.md), [gin-context.md](gin-context.md).
