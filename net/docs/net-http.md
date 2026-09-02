# 🌐 `net/http` Itu Sebenernya Apa Sih?

## TL;DR

`net/http` itu **package bawaan Go** (bukan library luar, bukan framework, nggak perlu `go get`) yang isinya semua peralatan buat bikin dan manggil HTTP server/client. Dia bagian dari **standard library** Go — sama kayak `fmt` atau `errors`, tinggal `import "net/http"` dan langsung bisa dipakai, nggak perlu install apapun.

```go
import "net/http" // udah ke-install otomatis bareng Go, gratis dipakai
```

## Kenapa Namanya `net/http`?

Ini soal struktur package standar library Go. `net` adalah package "induk" yang isinya semua urusan **networking** level rendah (TCP, UDP, DNS, IP, dll). Di dalam `net`, ada beberapa sub-package buat protokol spesifik:

```
net/          <- networking dasar (TCP, UDP, IP, dll)
net/http/     <- protokol HTTP, dibangun di ATAS net dasar
net/mail/     <- parsing email
net/url/      <- parsing URL
net/smtp/     <- protokol SMTP (kirim email)
```

Ini juga jawaban dari masalah kamu waktu itu (inget `module Net` yang bentrok?) — karena `net` itu nama package **root** yang penting banget di standard library, hati-hati jangan sampai nama module/package kamu sendiri nabrak nama ini.

## Apa Aja Isinya `net/http`?

Package ini se-"paket" isinya — nggak cuma buat server, tapi juga client:

### 1. Bikin HTTP Server
```go
http.HandleFunc("/", handler)          // daftarin route
http.ListenAndServe(":8080", nil)      // nyalain server, dengerin di port 8080
```

### 2. Bikin HTTP Client (Manggil API Lain)
```go
resp, err := http.Get("https://api.example.com/data")
```
Ini yang kepake kalau Go-mu justru jadi "pemanggil" API lain, bukan cuma "penyedia" API.

### 3. Tipe-tipe Penting yang Udah Kamu Pakai

| Tipe/Fungsi | Fungsinya |
|---|---|
| `http.ResponseWriter` | "kertas" buat nulis balik response ke client (parameter `w` di handler kamu) |
| `*http.Request` | data lengkap soal request yang masuk — method, header, body, URL, dll (parameter `r`) |
| `http.HandleFunc(pattern, handler)` | daftarin function jadi "penjaga" buat path tertentu |
| `http.Error(w, msg, code)` | shortcut kirim response error |
| `http.StatusOK`, dll | konstanta kode status HTTP (lihat [status-code.md](status-code.md)) |

## Kenapa Go Standard Library-nya Sekomplit Ini?

Ini salah satu **kelebihan besar** Go dibanding banyak bahasa lain. Di Node.js misalnya, bikin server HTTP dasar sebenernya juga bisa pakai `http` bawaan, tapi ekosistemnya lebih "mengarahkan" orang buat langsung pakai Express dari awal. Di Go, standard library-nya udah **cukup lengkap dan production-ready** buat bikin server HTTP beneran tanpa dependency luar sama sekali — makanya kita bisa latihan bikin REST API dari nol pakai `net/http` doang, tanpa install Gin/Echo/apapun.

Ini juga alasan kenapa framework kayak **Gin** itu bisa "ringan" — karena mereka nggak reinvent semua dari nol, cuma nambahin **lapisan kenyamanan** di atas `net/http` yang udah solid.

## Sejak Go 1.22: Makin Powerful

Dulu (sebelum Go 1.22), `net/http` agak "primitif" soal routing — nggak bisa bedain method (`GET`/`POST`) langsung di pattern, nggak ada path parameter bawaan (`/user/{id}`). Ini salah satu alasan besar orang pindah ke framework kayak Gin.

Sejak **Go 1.22** (2024), fitur-fitur itu ditambahin langsung ke `net/http`:
```go
http.HandleFunc("GET /profile", profilHandler)
http.HandleFunc("POST /user", createUserHandler)
http.HandleFunc("GET /user/{id}", getUserHandler) // r.PathValue("id")
```
Makanya sekarang makin banyak yang berani pakai `net/http` polos buat project kecil-menengah, tanpa buru-buru pasang framework.

## `net/http` vs Gin/Echo/Fiber — Ini Bukan "Lawan"

Penting dipahami: Gin **bukan pengganti** `net/http`, tapi **dibangun di atasnya**. Di balik layar, Gin tetap pakai `net/http` sebagai fondasi — dia cuma nambahin:
- Routing yang lebih ekspresif (grouping, middleware chaining).
- Binding otomatis JSON ke struct (kamu nulis manual pakai `json.NewDecoder(...).Decode(...)`, Gin sediain `c.BindJSON(...)`).
- Middleware ecosystem siap pakai (auth, logging, CORS, dll).

Jadi belajar `net/http` dulu itu **bukan buang waktu** — begitu nanti pindah ke Gin, semua konsep intinya (handler, request, response, status code) tetap sama persis, cuma syntax-nya lebih ringkes.

## Poin Penting

- `net/http` itu **standard library**, bukan framework luar — udah otomatis ada begitu Go ke-install, nggak butuh `go get`.
- Isinya lengkap buat 2 arah: jadi **server** (nerima request) dan jadi **client** (manggil API lain).
- Sejak Go 1.22, routing-nya udah cukup nyaman (method-based routing + path parameter) buat kebanyakan kebutuhan project kecil-menengah.
- Framework (Gin, dkk) itu "lapisan tambahan" di atas `net/http`, bukan pengganti total — fondasinya tetap sama.
