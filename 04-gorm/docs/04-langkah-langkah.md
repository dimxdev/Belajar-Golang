# ✅ Langkah-Langkah Bangun Project (Checklist)

Urutan yang disarankan buat bikin `04-gorm/` dari nol. Bangun **dari bawah ke atas** (database dulu, HTTP terakhir) — tiap langkah bisa dites sebelum lanjut.

---

## Fase 0: Persiapan

- [ ] Install PostgreSQL (atau MySQL) di lokal, atau jalanin lewat Docker
- [ ] Bikin database kosong: `CREATE DATABASE belajar_gorm;`
- [ ] `cd 04-gorm && go mod init belajar-gorm`
- [ ] `go get` semua dependency (lihat [02-setup-dan-koneksi.md](02-setup-dan-koneksi.md))
- [ ] Bikin `.env` + `.env.example`, tambahin `.env` ke `.gitignore` root

## Fase 1: Fondasi (Config + Database)

- [ ] `internal/config/config.go` — struct `Config` + `Load()`
- [ ] `internal/model/user.go` — struct `User` (biar `AutoMigrate` ada isinya)
- [ ] `internal/database/database.go` — `Connect()` + `Migrate()`
- [ ] `cmd/api/main.go` sementara: cuma `config.Load()` → `database.Connect()` → `database.Migrate()` → `log.Println("ok")`
- [ ] **Tes**: `go run ./cmd/api` — harus jalan tanpa error, dan cek di database tabel `users` udah kebentuk

## Fase 2: Repository

- [ ] `internal/repository/user_repository.go` — struct + `NewUserRepository` + `FindAll`, `FindByID`, `FindByEmail`, `Create`
- [ ] **Tes** (opsional, sementara di `main.go`): bikin 1 user manual lewat `repo.Create(...)`, terus `repo.FindAll()`, print hasilnya

## Fase 3: Service

- [ ] `internal/service/user_service.go` — struct + `NewUserService` + error domain (`ErrUserNotFound`, `ErrEmailTaken`) + `GetAll`, `GetByID`, `Create`
- [ ] Pindahin logic "cek email unik" dari mana pun ke sini

## Fase 4: Handler + Router

- [ ] `internal/model/user.go` — tambahin struct `CreateUserInput` dengan tag `binding`
- [ ] `internal/handler/user_handler.go` — struct + `NewUserHandler` + `GetAll`, `GetByID`, `Create`
- [ ] `internal/router/router.go` — `Setup(db)`: wiring + daftar route
- [ ] `cmd/api/main.go` final: tambahin `router.Setup(db)` + `r.Run(...)`
- [ ] **Tes**: `go run ./cmd/api`, lalu pakai Postman/curl

## Fase 5: Middleware (Opsional)

- [ ] `internal/middleware/logger.go` — custom logger
- [ ] Pasang di `router.go` lewat `r.Use(...)` atau per-group

---

## Cara Tes Endpoint (curl)

```bash
# Buat user
curl -X POST http://localhost:8000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Budi","email":"budi@mail.com","age":25}'

# List semua user
curl http://localhost:8000/api/v1/users

# Ambil user by ID
curl http://localhost:8000/api/v1/users/1

# Tes validasi (harus gagal — umur < 17)
curl -X POST http://localhost:8000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Ani","email":"ani@mail.com","age":15}'

# Tes email duplikat (harus 409)
curl -X POST http://localhost:8000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Budi 2","email":"budi@mail.com","age":30}'
```

---

## Jalanin PostgreSQL Cepat Pakai Docker (Kalau Belum Install)

```bash
docker run --name pg-belajar -e POSTGRES_PASSWORD=rahasia123 -e POSTGRES_DB=belajar_gorm -p 5432:5432 -d postgres:16
```

Sesuaikan `.env`:
```env
DB_PASSWORD=rahasia123
DB_NAME=belajar_gorm
```

---

## Kalau Udah Lancar, Latihan Lanjutan

1. Tambah fitur **Product** (ulangi pola yang sama: model → repo → service → handler)
2. Tambah endpoint `PUT /users/:id` dan `DELETE /users/:id`
3. Tambah **relasi** — `User` punya banyak `Order`, coba `Preload` (lihat [../../03-gin/docs/08-gorm.md](../../03-gin/docs/08-gorm.md))
4. Tambah **pagination** di `GetAll` (`?page=1&limit=10`) — pakai `c.DefaultQuery` + GORM `.Limit().Offset()`
5. Pindahin response format ke helper reusable (biar `gin.H{"status": ..., "data": ...}` nggak ditulis ulang di tiap handler)
6. Tambah Swagger (lihat catatan Swagger yang udah kamu pelajari) — dokumentasiin tiap endpoint

---

## Poin Penting

- Bangun **dari bawah ke atas**: config → database → model → repository → service → handler → router.
- Tes tiap fase sebelum lanjut — jangan tulis semua sekaligus baru run di akhir (susah nyari error-nya).
- Fase 1 selesai = tabel kebentuk di database. Itu milestone pertama.
- Fase 4 selesai = endpoint bisa dipanggil dari Postman. Itu MVP-nya.
- Fitur kedua (Product) tinggal copy pola fitur pertama — kalau kerasa "ngulang banget", berarti strukturnya udah bener.
