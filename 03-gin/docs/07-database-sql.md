# 🗄️ Koneksi ke Database (`database/sql`) — PostgreSQL / MySQL

Sampai sini, semua data di handler kamu masih **hardcode** (`Profile{Name: "Dimas"}`). Langkah berikutnya: ambil/simpan data ke **database beneran**. Go punya package standar `database/sql` buat ini.

---

## Konsep: `database/sql` + Driver

`database/sql` itu **cuma "interface umum"** — dia nyediain cara standar buat query, transaksi, connection pool, dll, tapi **nggak tau** cara ngobrol sama database spesifik (PostgreSQL beda protokol sama MySQL). Buat itu, butuh **driver** — library pihak ketiga yang "menerjemahkan" perintah `database/sql` ke bahasa database tertentu.

```
Kode kamu -> database/sql (standar) -> driver (spesifik) -> Database
```

| Database | Driver yang umum dipakai |
|---|---|
| PostgreSQL | `github.com/lib/pq` atau `github.com/jackc/pgx` |
| MySQL | `github.com/go-sql-driver/mysql` |

---

## 1. Install Driver

**PostgreSQL:**
```bash
go get github.com/lib/pq
```

**MySQL:**
```bash
go get github.com/go-sql-driver/mysql
```

---

## 2. Import Driver Pakai Blank Import (`_`)

Ini pola yang udah kita bahas di materi Swagger — **blank import** (`_`). Kamu nggak manggil function apapun dari driver secara langsung; driver cuma perlu di-import biar `func init()`-nya jalan dan **daftarin diri** ke `database/sql`.

```go
import (
	"database/sql"

	_ "github.com/lib/pq"              // PostgreSQL
	// _ "github.com/go-sql-driver/mysql" // MySQL (pilih salah satu)
)
```

---

## 3. Buka Koneksi (`sql.Open`)

```go
func ConnectDB() (*sql.DB, error) {
	// PostgreSQL
	dsn := "host=localhost port=5432 user=postgres password=rahasia dbname=belajar_go sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	// MySQL (formatnya beda):
	// dsn := "root:rahasia@tcp(localhost:3306)/belajar_go?parseTime=true"
	// db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	// sql.Open TIDAK langsung konek — cuma validasi argumen.
	// Ping() yang benar-benar nyoba buka koneksi ke database.
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
```

### DSN (Data Source Name) — String Koneksi

Ini "alamat + kredensial" database. Formatnya beda per driver:

**PostgreSQL (`lib/pq`):**
```
host=localhost port=5432 user=postgres password=rahasia dbname=belajar_go sslmode=disable
```

**MySQL (`go-sql-driver/mysql`):**
```
root:rahasia@tcp(localhost:3306)/belajar_go?parseTime=true
```

| Bagian | Arti |
|---|---|
| `host` / `tcp(...)` | alamat server database |
| `port` | `5432` (Postgres default), `3306` (MySQL default) |
| `user` / `password` | kredensial login |
| `dbname` / `/belajar_go` | nama database yang dipakai |
| `sslmode=disable` | matiin SSL (buat development lokal) |
| `parseTime=true` (MySQL) | biar kolom `DATETIME`/`TIMESTAMP` otomatis jadi `time.Time` Go |

### ⚠️ `sql.Open` vs `db.Ping`

- `sql.Open(...)` **nggak** langsung nyambung ke database — dia cuma validasi format DSN & siapin object `*sql.DB`. Jadi `sql.Open` hampir nggak pernah error meski database-nya mati.
- `db.Ping()` yang **beneran** nyoba buka koneksi. Selalu panggil `Ping()` setelah `Open()` buat mastiin koneksi valid dari awal.

---

## 4. `*sql.DB` Itu Connection Pool, Bukan 1 Koneksi

Ini konsep penting: object `*sql.DB` yang dibalikin `sql.Open` **bukan** 1 koneksi tunggal, tapi **pool** (kumpulan) koneksi yang dikelola otomatis oleh Go. Implikasinya:

1. **Bikin `*sql.DB` sekali aja** — biasanya di `main()` pas startup, terus dioper ke seluruh aplikasi (jangan `sql.Open` di tiap handler!).
2. **Aman dipakai concurrent** — banyak goroutine (misal banyak request HTTP bersamaan) bisa pakai `*sql.DB` yang sama, Go otomatis atur pembagian koneksinya.
3. **Nggak perlu buru-buru di-`Close()`** — karena dia hidup selama aplikasi jalan. `db.Close()` biasanya cuma dipanggil pas aplikasi mau shutdown.

Konfigurasi pool (opsional, tapi disarankan buat production):
```go
db.SetMaxOpenConns(25)                 // maksimal koneksi terbuka bersamaan
db.SetMaxIdleConns(25)                 // maksimal koneksi "nganggur" yang disimpan
db.SetConnMaxLifetime(5 * time.Minute) // umur maksimal 1 koneksi sebelum dibuang
```

---

## 5. Integrasi ke Project Gin

### Struktur: Simpan `*sql.DB` dan Oper ke Handler

Handler Gin signature-nya fixed (`func(c *gin.Context)`), jadi `*sql.DB` nggak bisa dilewatin lewat parameter. Solusi umum: pakai **closure** (pola pabrik yang sama kayak middleware).

```go
// db-nya "dititipkan" ke handler lewat closure
func GetProfilHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var profil Profile

		row := db.QueryRow("SELECT name, age FROM profiles WHERE id = $1", 1)
		if err := row.Scan(&profil.Name, &profil.Age); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(404, gin.H{"message": "profil tidak ditemukan"})
				return
			}
			c.JSON(500, gin.H{"message": "gagal ambil data"})
			return
		}

		c.JSON(200, profil)
	}
}
```

### `main.go`

```go
func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Gagal konek database:", err) // Fatal = print + os.Exit(1)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/profile", GetProfilHandler(db)) // db dioper ke handler
	router.Run(":8000")
}
```

---

## 6. Operasi Dasar

### Query 1 Baris — `QueryRow` + `Scan`

```go
var profil Profile
row := db.QueryRow("SELECT name, age FROM profiles WHERE id = $1", 1)
err := row.Scan(&profil.Name, &profil.Age) // isi field struct lewat pointer
```

`Scan(&profil.Name, &profil.Age)` — butuh **pointer** (`&`), sama alasannya kayak `json.Decode` / `c.ShouldBindJSON`: dia perlu **mengisi** variabel yang kamu kasih.

### Query Banyak Baris — `Query` + loop `rows.Next()`

```go
rows, err := db.Query("SELECT name, age FROM profiles")
if err != nil {
	return nil, err
}
defer rows.Close() // WAJIB, biar koneksi balik ke pool

var hasil []Profile
for rows.Next() {
	var p Profile
	if err := rows.Scan(&p.Name, &p.Age); err != nil {
		return nil, err
	}
	hasil = append(hasil, p)
}
```

### INSERT / UPDATE / DELETE — `Exec`

```go
result, err := db.Exec(
	"INSERT INTO profiles (name, age) VALUES ($1, $2)",
	"Budi", 25,
)
if err != nil {
	return err
}

rowsAffected, _ := result.RowsAffected() // berapa baris kena
// lastID, _ := result.LastInsertId()    // ID baris baru (MySQL; di Postgres pakai RETURNING)
```

---

## 7. Placeholder Parameter — JANGAN String Concatenation!

Perhatiin `$1`, `$2` (Postgres) atau `?` (MySQL) di query di atas — itu **placeholder**. Nilai dikirim **terpisah** dari string query.

**❌ JANGAN PERNAH kayak gini** (rawan SQL Injection):
```go
db.Query("SELECT * FROM users WHERE name = '" + namaInput + "'") // BAHAYA!
```

**✅ SELALU pakai placeholder:**
```go
db.Query("SELECT * FROM users WHERE name = $1", namaInput) // Postgres
db.Query("SELECT * FROM users WHERE name = ?", namaInput)  // MySQL
```

Dengan placeholder, driver + database yang urus "escaping" nilai input dengan aman — input jahat kayak `'; DROP TABLE users; --` diperlakukan sebagai **teks biasa**, bukan perintah SQL.

| Database | Gaya Placeholder |
|---|---|
| PostgreSQL | `$1`, `$2`, `$3`, ... (bernomor) |
| MySQL | `?`, `?`, `?`, ... (posisional) |

---

## 8. `defer rows.Close()` — Jangan Lupa

Tiap `db.Query(...)` (yang balikin `*sql.Rows`) **wajib** ditutup dengan `rows.Close()`, biasa pakai `defer` biar otomatis kepanggil pas function selesai:

```go
rows, err := db.Query("...")
if err != nil { return err }
defer rows.Close() // <- ini
```

Kalau lupa, koneksi itu **nggak balik** ke pool, dan lama-lama pool-nya habis → aplikasi macet nggak bisa query lagi. (Ini nggak berlaku buat `QueryRow` dan `Exec` — cuma `Query` yang perlu `Close()` manual.)

---

## Poin Penting

- `database/sql` = interface standar; **driver** (`lib/pq`, `go-sql-driver/mysql`) = penerjemah ke database spesifik, di-import pakai blank import (`_`).
- `sql.Open` nggak langsung konek — `db.Ping()` yang beneran nyoba koneksi, panggil setelah `Open`.
- `*sql.DB` itu **connection pool**, bikin **sekali** di `main()`, oper ke seluruh aplikasi (lewat closure buat handler Gin), aman dipakai concurrent.
- **Selalu** pakai placeholder (`$1` / `?`) buat nilai input — jangan string concatenation (SQL Injection).
- `db.Query` → butuh `defer rows.Close()`. `db.QueryRow` & `db.Exec` → nggak perlu.
- `Scan(&field)` butuh pointer, sama konsepnya kayak `Decode`/`ShouldBindJSON`.
- Cek `sql.ErrNoRows` khusus buat kasus "data nggak ketemu" (mirip pola `if !ada` yang dibahas sebelumnya) vs error lain (koneksi putus, query salah, dll).
- Untuk mapping struct ↔ tabel yang lebih otomatis (tanpa `Scan` manual field per field), nanti bisa lihat library seperti `sqlx`, `sqlc`, atau ORM seperti `GORM` — tapi paham `database/sql` dulu sebagai fondasi.
