# 📟 HTTP Status Code di Go (`net/http`)

Daftar konstanta status code yang paling sering dipakai di package `net/http`, beserta arti singkatnya. Dipakai lewat `w.WriteHeader(http.StatusXxx)` atau sebagai parameter ketiga di `http.Error(w, pesan, http.StatusXxx)`.

## 2xx — Sukses

| Konstanta | Kode | Arti |
|---|---|---|
| `http.StatusOK` | 200 | Request berhasil (default kalau `WriteHeader` gak dipanggil manual) |
| `http.StatusCreated` | 201 | Berhasil, dan resource baru berhasil dibuat (biasa dipake abis `POST`) |
| `http.StatusAccepted` | 202 | Request diterima, tapi diproses async/belum tentu selesai |
| `http.StatusNoContent` | 204 | Berhasil, tapi gak ada body buat dikirim balik (biasa abis `DELETE`) |

## 3xx — Redirect

| Konstanta | Kode | Arti |
|---|---|---|
| `http.StatusMovedPermanently` | 301 | Resource pindah permanen ke URL lain |
| `http.StatusFound` | 302 | Redirect sementara |
| `http.StatusNotModified` | 304 | Dipake buat caching, data di client masih sama/valid |

## 4xx — Error dari sisi client (request-nya yang salah)

| Konstanta | Kode | Arti |
|---|---|---|
| `http.StatusBadRequest` | 400 | Request salah format / body invalid (misal JSON rusak) |
| `http.StatusUnauthorized` | 401 | Belum login / gak ada token / token invalid |
| `http.StatusForbidden` | 403 | Udah login, tapi gak punya akses ke resource ini |
| `http.StatusNotFound` | 404 | Resource / endpoint yang diminta gak ketemu |
| `http.StatusMethodNotAllowed` | 405 | Method salah (misal endpoint cuma nerima `POST`, tapi di-`GET`) |
| `http.StatusConflict` | 409 | Konflik, misal data yang mau dibuat udah ada/duplikat |
| `http.StatusUnprocessableEntity` | 422 | Format request bener, tapi gagal validasi bisnis (misal harga < 0) |
| `http.StatusTooManyRequests` | 429 | Kena rate limit, request kebanyakan dalam waktu singkat |

## 5xx — Error dari sisi server

| Konstanta | Kode | Arti |
|---|---|---|
| `http.StatusInternalServerError` | 500 | Error gak terduga di server (misal bug, panic, dll) |
| `http.StatusNotImplemented` | 501 | Server belum support fitur/method yang diminta |
| `http.StatusBadGateway` | 502 | Server jadi proxy/gateway, tapi dapet response invalid dari server lain |
| `http.StatusServiceUnavailable` | 503 | Server lagi down / overload / maintenance |
| `http.StatusGatewayTimeout` | 504 | Server proxy/gateway kelamaan nunggu response dari server lain |

## Cara pakai

```go
// Set status doang, body ditulis manual setelahnya (bisa JSON, dll)
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(data)

// Shortcut khusus buat error, otomatis set status + kirim pesan teks
http.Error(w, "Data tidak valid", http.StatusBadRequest)
```

**Catatan:** yang paling sering dipakai sehari-hari di REST API sederhana biasanya cuma segelintir: `200`, `201`, `204`, `400`, `401`, `403`, `404`, `422`, `500`. Lainnya lebih jarang dipakai kecuali kasus spesifik (redirect, caching, proxy, dll).

Daftar lengkap semua konstanta status code bisa dicek di dokumentasi resmi: https://pkg.go.dev/net/http#pkg-constants
