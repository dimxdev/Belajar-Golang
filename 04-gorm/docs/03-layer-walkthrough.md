# 🧩 Walkthrough Tiap Layer (dengan Skeleton Kode)

Contoh kasus: fitur **User** (list, ambil by ID, buat baru). Dibangun dari bawah ke atas: model → repository → service → handler → router.

---

## Layer 1: `internal/model/user.go`

```go
package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
```

- Tag `json:"..."` → bentuk JSON pas dikirim ke client (dari [../../03-gin/docs/04-post.md](../../03-gin/docs/04-post.md)).
- Tag `gorm:"..."` → aturan kolom database (dari [../../03-gin/docs/08-gorm.md](../../03-gin/docs/08-gorm.md)).

### Struct Input Terpisah (Best Practice)

Jangan pakai struct `User` langsung buat nerima input dari client — bikin struct khusus, biar client nggak bisa "maksa" isi field kayak `ID` atau `CreatedAt`:

```go
package model

type CreateUserInput struct {
	Name  string `json:"name" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,gte=17"`
}
```

Tag `binding:"..."` → validasi (dari [../../03-gin/docs/05-validation.md](../../03-gin/docs/05-validation.md)).

---

## Layer 2: `internal/repository/user_repository.go`

**Satu-satunya** layer yang nyentuh `db`. Isinya CRUD murni, nggak ada logic bisnis.

```go
package repository

import (
	"belajar-gorm/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// constructor — bikin instance repository dengan db yang di-inject
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}
```

### Pola "Constructor" (`NewXxx`)

`NewUserRepository(db)` itu **bukan keyword Go** — cuma **konvensi** buat function yang "bikin dan siapkan" sebuah struct. Isinya struct `UserRepository` nyimpen `db` sebagai field, jadi semua method-nya (`FindAll`, `Create`, dll) bisa akses `r.db` tanpa perlu di-passing berulang. Ini pola **dependency injection** — `db` "disuntikkan" dari luar pas pembuatan, bukan dibikin di dalam.

---

## Layer 3: `internal/service/user_service.go`

Business logic. Manggil repository buat urusan data. **Nggak tau soal HTTP.**

```go
package service

import (
	"errors"

	"belajar-gorm/internal/model"
	"belajar-gorm/internal/repository"

	"gorm.io/gorm"
)

// error khusus domain — biar handler bisa bedain jenis kegagalan
var (
	ErrUserNotFound = errors.New("user tidak ditemukan")
	ErrEmailTaken   = errors.New("email sudah digunakan")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound // terjemahin error GORM jadi error domain
		}
		return nil, err
	}
	return user, nil
}

func (s *UserService) Create(input model.CreateUserInput) (*model.User, error) {
	// aturan bisnis: email harus unik
	existing, _ := s.repo.FindByEmail(input.Email)
	if existing != nil {
		return nil, ErrEmailTaken
	}

	user := &model.User{
		Name:  input.Name,
		Email: input.Email,
		Age:   input.Age,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
```

### Kenapa Bikin Error Sendiri (`ErrUserNotFound`, `ErrEmailTaken`)?

Service **nggak boleh** return `gorm.ErrRecordNotFound` mentah ke handler — itu bocorin detail "kita pakai GORM" ke layer atas. Sebagai gantinya, service nerjemahin ke error domain-nya sendiri (`ErrUserNotFound`), dan handler nanti map error itu ke status HTTP (`404`).

---

## Layer 4: `internal/handler/user_handler.go`

Lapisan HTTP. Baca request → panggil service → terjemahin hasil jadi response.

```go
package handler

import (
	"errors"
	"net/http"
	"strconv"

	"belajar-gorm/internal/model"
	"belajar-gorm/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal ambil data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // path param string -> int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id tidak valid"})
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func (h *UserHandler) Create(c *gin.Context) {
	var input model.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.service.Create(input)
	if err != nil {
		if errors.Is(err, service.ErrEmailTaken) {
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()}) // 409
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "gagal simpan"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": user})
}
```

Perhatiin: handler **nggak ada** logic "cek email unik" atau "umur minimal" — itu semua di service. Handler cuma: parse input, panggil service, `switch` jenis error → status code.

---

## Layer 5: `internal/router/router.go`

```go
package router

import (
	"belajar-gorm/internal/handler"
	"belajar-gorm/internal/repository"
	"belajar-gorm/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// --- Wiring: db -> repository -> service -> handler ---
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// --- Routes ---
	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.POST("", userHandler.Create)
		}
	}

	return r
}
```

### "Wiring" / Dependency Chain

Baris ini yang "merakit" semua layer:
```go
userRepo := repository.NewUserRepository(db)       // db     masuk ke repo
userService := service.NewUserService(userRepo)    // repo   masuk ke service
userHandler := handler.NewUserHandler(userService) // service masuk ke handler
```

Tiap layer nerima layer di bawahnya lewat constructor-nya. Ini bikin tiap layer **nggak tau** dari mana dependency-nya datang — gampang di-ganti pas testing (misal kasih `userRepo` palsu/mock).

---

## Ringkasan Alur Dependency

```
main.go
  └─ router.Setup(db)
       ├─ NewUserRepository(db)          ← nyentuh database
       │    └─ dipakai oleh ─────────┐
       ├─ NewUserService(userRepo) ◄──┘  ← business logic
       │    └─ dipakai oleh ─────────┐
       └─ NewUserHandler(userService) ◄┘ ← HTTP layer
            └─ didaftarkan ke route
```

---

## Poin Penting

- Tiap layer = 1 struct + constructor `NewXxx(dependency)` + method-method-nya.
- **Model**: struct data. Bikin struct input terpisah (`CreateUserInput`) buat nerima request, jangan pakai model DB langsung.
- **Repository**: cuma CRUD, satu-satunya yang pegang `db`.
- **Service**: aturan bisnis, bikin error domain sendiri (`ErrUserNotFound`), nggak tau HTTP.
- **Handler**: parse request → panggil service → map error ke status code. Nggak ada logic bisnis.
- **Router**: rakit semua (`db → repo → service → handler`), daftarin route.
- Pola `NewXxx(...)` = dependency injection — dependency disuntik dari luar, bukan dibikin di dalam. Bikin kode gampang di-test.
