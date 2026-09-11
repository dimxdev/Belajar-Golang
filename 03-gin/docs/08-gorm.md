# 🐘 Kenalan sama GORM (ORM buat Go)

Di materi [07-database-sql.md](07-database-sql.md), kamu belajar `database/sql` — nulis query SQL manual, `Scan` field satu-satu. GORM adalah **ORM** (Object-Relational Mapping) yang bikin itu jauh lebih ringkas: kamu kerja dengan **struct Go**, GORM yang nerjemahin ke SQL di belakang layar.

Kalau kamu pernah pakai **Prisma** (Node.js) atau **Eloquent** (Laravel/PHP), konsepnya mirip — "ngomong" ke database pakai kode/objek, bukan nulis SQL mentah.

---

## `database/sql` vs GORM — Bandingin Langsung

**Ambil 1 user pakai `database/sql`:**
```go
var u User
row := db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", 1)
err := row.Scan(&u.ID, &u.Name, &u.Age)
```

**Ambil 1 user pakai GORM:**
```go
var u User
err := db.First(&u, 1).Error
```

GORM otomatis: bikin query `SELECT * FROM users WHERE id = 1`, mapping hasilnya ke field struct `u`, handle `sql.ErrNoRows`, dll.

---

## 1. Install

```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres   # atau gorm.io/driver/mysql
```

Beda dari `database/sql`, driver GORM **bukan** blank import — dipakai langsung:
```go
import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)
```

---

## 2. Koneksi

```go
func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=rahasia dbname=belajar_go port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
```

DSN-nya sama persis kayak yang di [07-database-sql.md](07-database-sql.md). Yang beda cuma `gorm.Open(...)` gantiin `sql.Open(...)`, dan hasilnya `*gorm.DB` (bukan `*sql.DB`).

---

## 3. Model — Struct = Tabel

Di GORM, struct Go **merepresentasikan tabel**. Ini mirip "model" di Prisma schema atau Eloquent Model.

```go
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```

### Konvensi Otomatis GORM

GORM banyak "nebak" berdasarkan konvensi (biar kamu nulis lebih sedikit):

| Yang kamu tulis | Yang GORM asumsikan |
|---|---|
| `type User struct` | nama tabel: `users` (plural, huruf kecil) |
| Field `ID` (`uint`) | primary key, auto-increment |
| Field `CreatedAt` (`time.Time`) | otomatis diisi waktu row dibuat |
| Field `UpdatedAt` (`time.Time`) | otomatis di-update tiap row diubah |
| Field `Name string` | kolom `name` (snake_case) |

### Tag `gorm:"..."`

Mirip tag `json:"..."` dan `binding:"..."` yang udah kamu pelajari — ini instruksi khusus buat GORM:

```go
Email string `gorm:"unique"`              // kolom UNIQUE
Name  string `gorm:"not null;default:'anon'"` // NOT NULL + default value
Age   int    `gorm:"column:umur"`         // nama kolom custom (bukan "age")
Bio   string `gorm:"type:text"`           // tipe kolom spesifik
```

### `gorm.Model` — Shortcut Field Umum

GORM nyediain struct siap pakai buat field yang hampir selalu ada:
```go
type User struct {
	gorm.Model        // embed: otomatis dapat ID, CreatedAt, UpdatedAt, DeletedAt
	Name  string
	Email string
}
```
`gorm.Model` isinya `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` — pola struct embedding yang udah dibahas di [03-struct.md](../../01-basic/docs/03-struct.md).

---

## 4. Auto Migration — Bikin/Update Tabel dari Struct

```go
db.AutoMigrate(&User{}, &Product{})
```

GORM otomatis bikin tabel `users` & `products` sesuai struktur struct, atau **update** tabel yang udah ada kalau ada field baru. Ini mirip `prisma migrate` / `php artisan migrate`, cuma lebih "otomatis" (dan lebih terbatas — dia nggak hapus kolom/nggak handle perubahan kompleks; buat production biasanya tetap pakai migration tool terpisah).

Biasa dipanggil sekali di `main()` pas startup:
```go
func main() {
	db, _ := ConnectDB()
	db.AutoMigrate(&User{})
	// ...
}
```

---

## 5. CRUD Dasar

### Create

```go
user := User{Name: "Budi", Email: "budi@mail.com", Age: 25}
result := db.Create(&user)

if result.Error != nil {
	// handle error
}
fmt.Println(user.ID) // GORM otomatis isi ID setelah insert
```

### Read

```go
// Ambil 1 by primary key
var user User
db.First(&user, 1) // SELECT * FROM users WHERE id = 1 LIMIT 1

// Ambil 1 by kondisi
db.First(&user, "email = ?", "budi@mail.com")

// Ambil semua
var users []User
db.Find(&users) // SELECT * FROM users

// Ambil dengan kondisi
db.Where("age > ?", 20).Find(&users)

// Chaining (mirip query builder Eloquent/Prisma)
db.Where("age > ?", 20).Order("name asc").Limit(10).Find(&users)
```

### Update

```go
var user User
db.First(&user, 1)

user.Age = 30
db.Save(&user) // UPDATE semua field

// Atau update field tertentu aja
db.Model(&user).Update("age", 30)
db.Model(&user).Updates(User{Name: "Budi Baru", Age: 31}) // update beberapa field
```

### Delete

```go
db.Delete(&user, 1) // DELETE FROM users WHERE id = 1
```

⚠️ Kalau model kamu punya field `DeletedAt` (atau embed `gorm.Model`), `Delete` itu **soft delete** — row-nya nggak beneran dihapus, cuma diisi timestamp `deleted_at`, dan otomatis di-exclude dari query berikutnya. Buat hapus permanen: `db.Unscoped().Delete(&user, 1)`.

---

## 6. Integrasi ke Gin

Sama pola-nya kayak `database/sql` — oper `*gorm.DB` ke handler lewat closure:

```go
func GetUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(500, gin.H{"message": "gagal ambil data"})
			return
		}
		c.JSON(200, gin.H{"status": "success", "data": users})
	}
}

func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"message": "data tidak valid"})
			return
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(500, gin.H{"message": "gagal simpan"})
			return
		}

		c.JSON(201, gin.H{"status": "success", "data": user})
	}
}
```

`main.go`:
```go
func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

	router := gin.Default()
	router.GET("/users", GetUsersHandler(db))
	router.POST("/users", CreateUserHandler(db))
	router.Run(":8000")
}
```

---

## 7. Error Handling di GORM

GORM nggak return `error` langsung dari method-nya — dia nyimpen di field `.Error` dari hasil chain:

```go
result := db.First(&user, 1)
if result.Error != nil {
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"message": "user tidak ditemukan"})
		return
	}
	c.JSON(500, gin.H{"message": "error database"})
	return
}
```

`gorm.ErrRecordNotFound` itu padanan `sql.ErrNoRows` di `database/sql` — buat bedain "data emang nggak ada" vs "ada masalah lain".

---

## 8. Relasi Antar Tabel (Sekilas)

GORM bisa handle relasi (one-to-many, many-to-many) lewat struktur struct:

```go
type User struct {
	gorm.Model
	Name   string
	Orders []Order // one-to-many: 1 user punya banyak order
}

type Order struct {
	gorm.Model
	Total  float64
	UserID uint // foreign key ke User
}
```

Ambil user beserta order-nya (mirip `include` di Prisma / `with()` di Eloquent):
```go
var user User
db.Preload("Orders").First(&user, 1) // user + semua order-nya sekaligus
```

---

## GORM vs `database/sql` — Kapan Pakai yang Mana?

| | `database/sql` | GORM |
|---|---|---|
| Kontrol query | Penuh (kamu tulis SQL sendiri) | GORM yang generate (bisa di-override kalau perlu) |
| Boilerplate | Banyak (`Scan` manual, dll) | Sedikit |
| Kurva belajar | SQL murni | Perlu hafal API GORM |
| Performa | Sedikit lebih cepat (nggak ada layer abstraksi) | Sedikit overhead, tapi biasanya nggak signifikan |
| Query kompleks (join banyak, agregasi) | Lebih enak & transparan | Kadang ribet / tetap harus turun ke raw SQL |

**Rekomendasi belajar:** pahami [07-database-sql.md](07-database-sql.md) dulu (biar ngerti apa yang GORM lakuin di belakang layar), baru pindah ke GORM buat produktivitas sehari-hari. Kebanyakan project Go modern pakai GORM (atau `sqlx`/`sqlc` sebagai jalan tengah).

---

## Poin Penting

- GORM = ORM buat Go, konsepnya mirip Prisma (Node) / Eloquent (Laravel) — kerja dengan struct, bukan SQL mentah.
- Struct = tabel; tag `gorm:"..."` = instruksi kolom (mirip tag `json`/`binding` yang udah kamu kenal).
- GORM banyak nebak lewat **konvensi**: `User` → tabel `users`, field `ID` → primary key, `CreatedAt`/`UpdatedAt` → auto-timestamp.
- `AutoMigrate(&User{})` bikin/update tabel dari struct — praktis buat development.
- CRUD: `Create`, `First`/`Find`, `Save`/`Updates`, `Delete` — chaining `.Where().Order().Limit()` mirip query builder.
- Error dicek lewat `.Error` (bukan return value), `gorm.ErrRecordNotFound` = padanan `sql.ErrNoRows`.
- Integrasi Gin: oper `*gorm.DB` ke handler lewat closure, sama pola-nya kayak `*sql.DB`.
- `Delete` = soft delete kalau ada field `DeletedAt` — hati-hati, row nggak beneran hilang.
