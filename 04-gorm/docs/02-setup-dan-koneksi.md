# ⚙️ Setup Project & Koneksi Database

## 1. Inisialisasi Module

```bash
cd 04-gorm
go mod init belajar-gorm
```

## 2. Install Dependency

```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres        # atau gorm.io/driver/mysql
go get github.com/joho/godotenv       # buat baca file .env
```

---

## 3. File `.env` dan `.env.example`

**`.env`** (nilai asli — masukkan ke `.gitignore`, JANGAN di-commit):
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=rahasia123
DB_NAME=belajar_gorm
DB_SSLMODE=disable

SERVER_PORT=8000
```

**`.env.example`** (template — di-commit, tanpa nilai rahasia, biar orang lain tau env apa aja yang dibutuhin):
```env
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_SSLMODE=disable

SERVER_PORT=8000
```

> Jangan lupa tambahin `.env` ke `.gitignore` di root repo. Password database bocor di git history itu masalah keamanan serius.

---

## 4. `internal/config/config.go` — Baca Env Var

```go
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

func Load() *Config {
	// godotenv.Load() baca file .env dan masukin ke environment.
	// Gak fatal kalau file .env gak ada (misal di production, env var di-set langsung).
	_ = godotenv.Load()

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "belajar_gorm"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		ServerPort: getEnv("SERVER_PORT", "8000"),
	}
}

// getEnv: ambil env var, kalau kosong pakai nilai default
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
```

---

## 5. `internal/database/database.go` — Koneksi GORM

```go
package database

import (
	"fmt"

	"belajar-gorm/internal/config"
	"belajar-gorm/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gagal konek database: %w", err)
	}

	return db, nil
}

// Migrate: bikin/update tabel dari struct model
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Product{},
	)
}
```

- `fmt.Sprintf(...)` dipakai buat "merangkai" DSN dari potongan-potongan config.
- `fmt.Errorf("...: %w", err)` — **error wrapping** (dibahas di [../../01-basic/docs/08-error-handling.md](../../01-basic/docs/08-error-handling.md)) — bungkus error asli sambil nambah konteks.

---

## 6. Sambungkan di `cmd/api/main.go`

```go
package main

import (
	"log"

	"belajar-gorm/internal/config"
	"belajar-gorm/internal/database"
	"belajar-gorm/internal/router"
)

func main() {
	// 1. Load config
	cfg := config.Load()

	// 2. Konek database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err) // log.Fatal = print error + os.Exit(1)
	}

	// 3. Migrate tabel
	if err := database.Migrate(db); err != nil {
		log.Fatal("gagal migrate:", err)
	}

	// 4. Setup router (semua route + inject db)
	r := router.Setup(db)

	// 5. Jalanin server
	log.Printf("Server jalan di http://localhost:%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
```

Lihat betapa tipisnya `main.go` — cuma urutan langkah startup, nggak ada logic apa-apa. Ini yang diinginkan.

---

## Poin Penting

- `go mod init belajar-gorm` → nama module ini jadi prefix semua import internal (`belajar-gorm/internal/...`).
- Config **selalu** lewat env var + `.env`, jangan hardcode. `.env` masuk `.gitignore`, `.env.example` di-commit sebagai template.
- `database.Connect` bikin `*gorm.DB` sekali; objek itu dioper ke seluruh aplikasi lewat parameter (`router.Setup(db)` → diteruskan ke repository).
- `AutoMigrate` dipanggil sekali di startup, setelah koneksi berhasil.
- `main.go` = wiring doang. Kalau `main.go` kamu mulai panjang & ada `if` logic, berarti ada yang salah taruh.
- Lanjut: [03-layer-walkthrough.md](03-layer-walkthrough.md).
