# 🧱 Middleware di Gin (Logging & CORS)

Middleware adalah function yang dijalankan **di antara** request masuk dan handler-nya diproses — kayak "penjaga pos" yang bisa ngecek/ngubah/nambahin sesuatu sebelum (atau sesudah) request itu sampai ke handler tujuan.

```
Request masuk -> [Middleware 1] -> [Middleware 2] -> Handler -> Response keluar
```

Contoh kasus umum yang dikerjain middleware: catat log tiap request, cek token login (auth), izinin/tolak request dari domain lain (CORS), ukur berapa lama request diproses, dll.

---

## Middleware yang Udah Otomatis Aktif: `gin.Default()`

Inget dari [gin.md](gin.md), `gin.Default()` udah otomatis pasang **2 middleware bawaan**:

```go
router := gin.Default()
```

1. **Logger** — nyatet tiap request yang masuk ke terminal (method, path, status code, durasi proses).
2. **Recovery** — nangkep `panic` di dalam handler, biar server nggak langsung crash total, tapi balikin response `500` ke client.

Kalau pakai `gin.New()` (bukan `Default()`), dua middleware ini **nggak** otomatis ada — kamu mulai dari router polos tanpa middleware apapun.

```go
router := gin.New() // kosong, gak ada logger/recovery otomatis
```

### Contoh Log dari `gin.Default()`

Begitu server jalan dan ada request masuk, otomatis muncul di terminal semacam ini:
```
[GIN] 2026/09/06 - 10:15:23 | 200 |      1.234ms |       127.0.0.1 | GET      "/profile"
```
Ini yang bikin kamu nggak perlu nulis manual `fmt.Println` tiap ada request masuk — Gin udah handle otomatis lewat middleware logger bawaan itu.

---

## Bikin Middleware Sendiri (Custom)

Middleware di Gin itu cuma **function biasa** yang nerima `gin.HandlerFunc` (sama persis kayak handler biasa), bedanya dia manggil `c.Next()` di dalamnya buat "melanjutkan" ke middleware/handler berikutnya.

```go
func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		fmt.Println("Request masuk:", c.Request.Method, c.Request.URL.Path)

		c.Next() // lanjut ke middleware/handler berikutnya

		durasi := time.Since(start)
		fmt.Println("Selesai dalam:", durasi)
	}
}
```

### Cara Pasang Middleware

**Global** (berlaku ke SEMUA route):
```go
router := gin.New()
router.Use(MyLogger())
```

**Per-route** (cuma berlaku ke route tertentu):
```go
router.GET("/profile", MyLogger(), profilHandler)
```

**Per-group** (berlaku ke sekumpulan route yang di-grouping):
```go
admin := router.Group("/admin")
admin.Use(MyLogger())
{
	admin.GET("/dashboard", dashboardHandler)
	admin.GET("/users", usersHandler)
}
```

---

## 3 Cara Nulis Middleware: Simple vs Pabrik vs Anonymous Inline

Middleware bisa ditulis dengan beberapa gaya, tergantung kebutuhan. Ketiganya **sama-sama valid** secara teknis — bedanya soal reusability dan konsistensi kode.

### 1. Simple — Function Biasa Tanpa Parameter

```go
func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(401, gin.H{"error": "Token tidak ada"})
		c.Abort() // STOP, JANGAN lanjut ke handler
		return
	}

	c.Next() // token ada, boleh lanjut
}

func main() {
	router := gin.Default()
	router.GET("/profile", authMiddleware, profilHandler) // middleware KHUSUS route ini
	router.Run(":8000")
}
```

Ini middleware paling gampang — signature-nya `func(c *gin.Context)` langsung cocok sama `gin.HandlerFunc`, jadi bisa langsung dipasang tanpa dipanggil dulu (perhatiin: `authMiddleware`, **bukan** `authMiddleware()`). Cocok kalau logic-nya **fixed**, nggak ada nilai yang perlu beda-beda tiap dipakai di route berbeda.

### 2. Pabrik (Factory) — Bisa Disetel dengan Parameter

```go
func RequireAPIKey(validKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")

		if key != validKey { // validKey "nempel" di sini lewat closure
			c.JSON(401, gin.H{"message": "API key salah/tidak ada"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.GET("/partner-a/data", RequireAPIKey("kunci-rahasia-a"), dataHandler)
	router.GET("/partner-b/data", RequireAPIKey("kunci-rahasia-b"), dataHandler)
	router.Run(":8000")
}
```

`RequireAPIKey(validKey string)` itu bukan middleware-nya sendiri — dia "pabrik" yang **menghasilkan** middleware (function di dalamnya) sambil menitipkan `validKey` lewat closure. Dipanggil dengan kurung (`RequireAPIKey("kunci-rahasia-a")`, **ada argumennya**), hasilnya baru didaftarkan ke route. Cocok kalau logic-nya sama, tapi butuh **nilai berbeda-beda** tiap route (beda API key, beda role, beda limit, dll) — satu function dipakai berkali-kali tanpa duplikasi kode.

**Contoh lain pola pabrik — parameternya bisa lebih dari 1 nilai (variadic):**

```go
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetHeader("X-User-Role") // asumsi role dikirim di header

		for _, role := range allowedRoles {
			if role == userRole {
				c.Next() // role cocok, boleh lanjut
				return
			}
		}

		c.JSON(403, gin.H{"message": "Kamu tidak punya akses"})
		c.Abort()
	}
}

func main() {
	router := gin.Default()

	router.GET("/admin/dashboard", RequireRole("admin"), dashboardHandler)       // cuma admin
	router.GET("/reports", RequireRole("admin", "staff"), reportsHandler)         // admin ATAU staff

	router.Run(":8000")
}
```

Bedanya sama `RequireAPIKey`: parameternya `...string` (variadic), jadi 1 middleware ini bisa dipanggil dengan **1 atau lebih** role sekaligus — `RequireRole("admin")` buat 1 role, `RequireRole("admin", "staff")` buat beberapa role dalam satu pemanggilan. Closure-nya tetap sama konsepnya: `allowedRoles` (slice hasil dari variadic) "nempel" di dalam function yang di-return, dipakai buat loop cek kecocokan role tiap ada request masuk.

### 3. Anonymous Inline — Tanpa Nama, Tanpa Pabrik

```go
func main() {
	router := gin.Default()

	validKey := "kunci-rahasia-a" // variabel biasa di scope main()

	router.GET("/data", func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key != validKey { // langsung "nangkep" validKey dari scope luar
			c.JSON(401, gin.H{"message": "API key salah/tidak ada"})
			c.Abort()
			return
		}
		c.Next()
	}, dataHandler)

	router.Run(":8000")
}
```

Middleware-nya ditulis **langsung di tempat**, tanpa nama, tanpa function pembungkus terpisah. Tetap bisa akses `validKey` dari scope sekitarnya lewat closure, tanpa perlu di-passing lewat parameter. Cocok kalau middleware itu **cuma dipakai sekali**, di 1 tempat doang — kalau dipakai berkali-kali dengan nilai berbeda, jadi harus copy-paste blok kode yang mirip di tiap route (kurang reusable dibanding pola pabrik).

### Perbandingan Ketiganya

| Pola | Butuh parameter? | Reusable dgn nilai beda? | Cara pasang |
|---|---|---|---|
| **Simple** | Tidak | Tidak relevan (logic fixed) | `router.GET(path, authMiddleware, handler)` |
| **Pabrik** | Ya | Ya, tinggal panggil ulang beda argumen | `router.GET(path, RequireAPIKey("x"), handler)` |
| **Anonymous inline** | Bisa "nyontek" dari scope luar | Tidak, tiap tempat harus nulis ulang | `router.GET(path, func(c *gin.Context) {...}, handler)` |

### Standar di Dunia Profesional

Di kode production, **pola pabrik** yang paling umum jadi konvensi/standar — bahkan dipakai juga buat middleware yang belum butuh parameter sekarang, contohnya middleware bawaan Gin sendiri:
```go
func Recovery() gin.HandlerFunc { // tetap pakai () -> gin.HandlerFunc walau tanpa parameter
	return func(c *gin.Context) { ... }
}
```
dipakai lewat `router.Use(gin.Recovery())` (ada kurungnya). Alasannya: **konsisten** (semua middleware "kebentuk" sama di seluruh codebase), dan **future-proof** (kalau nanti butuh nambah parameter, cara pemanggilannya `router.Use(Recovery())` tetap sama, tinggal isi argumen). Tapi ini konvensi tim, bukan aturan mutlak Go/Gin — pola simple maupun anonymous inline tetap valid dipakai kalau memang sesuai kebutuhannya.

---

## `c.Next()` dan `c.Abort()`

Dua method ini kunci buat ngerti alur middleware:

### `c.Next()` — Lanjut ke Berikutnya

```go
func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Sebelum handler")
		c.Next() // jalanin middleware/handler selanjutnya dulu
		fmt.Println("Sesudah handler") // baru lanjut ke sini SETELAH semua di bawahnya selesai
	}
}
```
Kode **sebelum** `c.Next()` jalan duluan (fase "masuk"), kode **sesudah** `c.Next()` jalan belakangan setelah handler tujuan selesai diproses (fase "keluar"). Ini yang bikin middleware bisa ngukur durasi kayak contoh logger di atas.

### `c.Abort()` — Berhenti, Jangan Lanjut

```go
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(401, gin.H{"message": "Token tidak ditemukan"})
			c.Abort() // STOP di sini, JANGAN lanjut ke handler
			return
		}

		c.Next() // token ada, boleh lanjut
	}
}
```
Kalau `c.Abort()` dipanggil, semua middleware/handler **setelahnya** di rantai itu **dilewati** — request langsung berhenti dengan response yang udah kamu kirim (misal `401` di atas).

---

## CORS (Cross-Origin Resource Sharing)

### Apa itu CORS dan Kenapa Penting?

Browser punya aturan keamanan: sebuah halaman web dari domain `A` **nggak boleh** manggil API di domain `B` secara default, kecuali server `B` secara eksplisit **mengizinkan**. Ini yang disebut **CORS policy**.

Contoh kasus nyata: frontend kamu jalan di `http://localhost:3000` (React/Vue), backend Gin kamu jalan di `http://localhost:8000`. Kalau backend nggak diatur CORS-nya, request dari frontend bakal **ditolak browser** dengan error semacam:
```
Access to fetch at 'http://localhost:8000/profile' from origin 'http://localhost:3000'
has been blocked by CORS policy
```

### Solusi 1: Middleware CORS Manual

```go
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // izinin semua domain
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // browser suka kirim OPTIONS dulu buat "cek izin" (preflight)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/profile", profilHandler)
	router.Run(":8000")
}
```

### Solusi 2: Pakai Library `gin-contrib/cors` (Lebih Praktis)

```bash
go get github.com/gin-contrib/cors
```

```go
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // domain yang diizinkan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/profile", profilHandler)
	router.Run(":8000")
}
```

Library ini lebih disarankan buat project beneran, soalnya udah nangani banyak edge case CORS (preflight request, credential, dll) yang kalau ditulis manual gampang kelewatan detailnya.

### Penjelasan Config CORS

| Config | Fungsi |
|---|---|
| `AllowOrigins` | domain mana aja yang boleh akses API ini (`*` = semua domain, tapi hati-hati kalau `AllowCredentials: true`) |
| `AllowMethods` | method HTTP mana aja yang diizinkan (`GET`, `POST`, dll) |
| `AllowHeaders` | header custom mana yang boleh dikirim client (misal `Authorization` buat token) |
| `AllowCredentials` | boleh nggak client kirim cookie/credential bareng request |

⚠️ **Penting:** `AllowOrigins: []string{"*"}` (semua domain boleh) itu **nggak boleh** dipakai bareng `AllowCredentials: true` — browser bakal nolak kombinasi ini karena alasan keamanan. Kalau butuh credential, `AllowOrigins` harus disebut domain spesifik, bukan `*`.

---

## Contoh Gabungan: Logger + CORS + Auth

```go
func main() {
	router := gin.Default() // logger + recovery udah otomatis

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	router.GET("/profile", profilHandler) // publik, gak perlu auth

	admin := router.Group("/admin")
	admin.Use(AuthMiddleware()) // cuma route di dalam group ini yang kena AuthMiddleware
	{
		admin.GET("/dashboard", dashboardHandler)
	}

	router.Run(":8000")
}
```

---

## Poin Penting

- Middleware itu function `gin.HandlerFunc` biasa yang manggil `c.Next()` buat lanjut ke tahap berikutnya — tanpa `c.Next()`, request bakal "macet" di situ.
- `c.Abort()` dipakai buat menghentikan rantai request lebih awal (misal auth gagal), tanpa lanjut ke handler tujuan.
- `router.Use(...)` pasang middleware **global**; taruh middleware sebagai argumen tambahan di `router.GET(path, middleware, handler)` buat middleware **per-route**; `group.Use(...)` buat middleware **per-group**.
- CORS wajib diatur kalau frontend & backend kamu beda origin (domain/port beda) — tanpa ini, browser bakal blokir request dari frontend meskipun API-nya sebenarnya jalan normal.
- Untuk CORS di project serius, lebih disarankan pakai `gin-contrib/cors` daripada nulis manual, biar nggak kelewatan edge case.
- Referensi terkait: [gin.md](gin.md), [gin-context.md](gin-context.md).
