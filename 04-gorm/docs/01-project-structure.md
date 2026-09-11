# 🏗️ Struktur Project Profesional (Gin + GORM)

Folder `04-gorm/` ini fokus latihan bikin REST API dengan **struktur berlapis (layered architecture)** yang umum dipakai di project Go production — beda dari `03-gin/` yang semua handler ditumpuk di folder `materi/`.

Materi kenalan GORM ada di [../../03-gin/docs/08-gorm.md](../../03-gin/docs/08-gorm.md) — baca dulu itu kalau belum.

---

## Struktur Folder

```
04-gorm/
├── cmd/
│   └── api/
│       └── main.go              # entry point — cuma wiring, gak ada logic
├── internal/
│   ├── config/
│   │   └── config.go            # baca env var (DB_HOST, DB_PASS, PORT, dll)
│   ├── database/
│   │   └── database.go          # koneksi GORM + AutoMigrate
│   ├── model/
│   │   ├── user.go              # struct User (= tabel users)
│   │   └── product.go
│   ├── repository/
│   │   ├── user_repository.go   # SATU-SATUNYA layer yang nyentuh db
│   │   └── product_repository.go
│   ├── service/
│   │   ├── user_service.go      # business logic (validasi, aturan, dll)
│   │   └── product_service.go
│   ├── handler/
│   │   ├── user_handler.go      # terima request, panggil service, kirim response
│   │   └── product_handler.go
│   ├── router/
│   │   └── router.go            # daftar semua route + middleware
│   └── middleware/
│       └── logger.go
├── .env                          # nilai config asli (JANGAN di-commit)
├── .env.example                  # template config (di-commit, tanpa nilai rahasia)
├── go.mod
└── README.md
```

---

## Tanggung Jawab Tiap Layer

Aturan emas: **tiap layer cuma boleh ngobrol sama layer tepat di bawahnya.** Handler nggak boleh langsung nyentuh database, service nggak boleh tau soal HTTP.

### `cmd/api/main.go` — Entry Point
Cuma "menyalakan mesin": baca config → konek DB → setup router → jalanin server. **Nggak ada logic bisnis di sini.** Idealnya cuma belasan baris.

### `internal/config` — Konfigurasi
Baca environment variable (`.env`) jadi struct Go yang gampang dipakai. Semua nilai yang beda antar environment (dev/staging/prod) — host database, password, port, secret — lewat sini, **bukan hardcode** di kode.

### `internal/database` — Koneksi Database
Bikin `*gorm.DB` (connection pool), jalanin `AutoMigrate`. Dibikin **sekali** di startup, terus objek `*gorm.DB`-nya dioper ke repository.

### `internal/model` — Representasi Data
Struct yang mewakili tabel database (`type User struct { ... }` dengan tag `gorm:"..."`). Cuma definisi data, nggak ada method logic berat. Dipakai bareng-bareng oleh repository, service, dan handler.

### `internal/repository` — Akses Database
**Satu-satunya** layer yang boleh manggil `db.Find()`, `db.Create()`, dll. Isinya function CRUD murni: `GetByID(id)`, `Create(user)`, `Update(user)`, `Delete(id)`. Nggak ada validasi/aturan bisnis — cuma "ambilin/simpen data".

> **Kenapa dipisah?** Kalau suatu saat ganti dari GORM ke `database/sql`, atau ganti PostgreSQL ke MongoDB — **cuma folder ini** yang perlu diubah. Handler & service nggak kesentuh.

### `internal/service` — Business Logic
Otak aplikasi. Contoh: "umur harus ≥ 17", "email harus unik, cek dulu ke repository", "kalau daftar berhasil, kirim email welcome". Service manggil repository buat urusan data, tapi **nggak tau apa-apa soal HTTP** (`gin.Context`, status code, dll).

### `internal/handler` — Lapisan HTTP
"Penerjemah" antara dunia HTTP dan service. Tugasnya: baca request (`c.ShouldBindJSON`, `c.Param`), panggil service, terjemahin hasilnya jadi response (`c.JSON` dengan status code yang sesuai). **Nggak ada logic bisnis** — kalau ada `if` yang mikirin aturan, itu harusnya di service.

### `internal/router` — Pendaftaran Route
Kumpulin semua `router.GET/POST/...` di satu tempat, pasang middleware. Biar `main.go` tetap bersih.

### `internal/middleware` — Middleware
Logger, auth, CORS, dll — konsepnya sama persis kayak [../../03-gin/docs/06-middleware.md](../../03-gin/docs/06-middleware.md).

---

## Kenapa Ada Folder `internal/`?

Inget dari pembahasan sebelumnya — `internal/` itu **nama spesial** yang dikenali Go compiler: apapun di dalamnya **cuma bisa di-import dari dalam module `04-gorm` sendiri**. Nggak bisa diakses project lain. Cocok buat kode aplikasi yang emang nggak dimaksudkan jadi library publik.

---

## Alur Data 1 Request (Contoh: `POST /users`)

```
1. Client kirim  POST /users  { "name": "Budi", "email": "budi@mail.com", "age": 25 }
                          │
2. router.go      ──────► arahkan ke  handler.CreateUser
                          │
3. handler        ──────► c.ShouldBindJSON(&input)   (baca + validasi format JSON)
                          panggil  service.CreateUser(input)
                          │
4. service        ──────► cek aturan bisnis: umur >= 17?
                          cek email udah dipakai?  ──► repository.GetByEmail(email)
                          kalau lolos:              ──► repository.Create(user)
                          │
5. repository     ──────► db.Create(&user)          (GORM → SQL INSERT)
                          │
6. (balik ke atas)        service return user       ──►
                          handler: c.JSON(201, ...)  ──►
7. Client terima  201 Created  { "status": "success", "data": { "id": 1, ... } }
```

Perhatiin: **error di layer manapun** (validasi gagal di service, DB error di repository) di-return ke atas sebagai `error`, dan **handler yang paling akhir** menerjemahkannya jadi status code HTTP yang sesuai (`400`, `422`, `500`, dll).

---

## Poin Penting

- **Pemisahan tanggung jawab (separation of concerns)**: handler = HTTP, service = logic, repository = database. Nggak boleh nyampur.
- **Ketergantungan mengalir ke bawah**: `handler → service → repository → database`. Nggak boleh kebalik.
- `internal/` bikin kode aplikasi nggak bisa di-import project luar.
- `cmd/api/main.go` cuma wiring — tipis, gampang dibaca.
- Config lewat env var, bukan hardcode — beda environment tinggal ganti `.env`.
- Struktur ini kelihatan "berlebihan" buat 1-2 endpoint, tapi kebayar pas project tumbuh: gampang di-test, gampang ganti teknologi, gampang kerja bareng tim.
- Langkah build ada di [04-langkah-langkah.md](04-langkah-langkah.md).
