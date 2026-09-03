# 📥 GET di Gin

Sama seperti [GET.md net/docs](../../net/docs/GET.md): method `GET` dipakai buat **mengambil data**, tanpa body request. Data tambahan lewat query parameter atau path parameter.

## Contoh Paling Dasar

```go
type Profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func profilHandler(c *gin.Context) {
	data := Profile{Name: "Dimas", Age: 20}
	c.JSON(200, data)
}
```

## Registrasi Route

```go
router.GET("/profile", profilHandler)
```
Beda dari `net/http` yang perlu tulis method di string pattern (`"GET /profile"`), di Gin method-nya langsung jadi nama function router.

## Baca Query Parameter (`?key=value`)

Contoh: `GET /search?nama=budi&limit=10`

```go
func searchHandler(c *gin.Context) {
	nama := c.Query("nama")               // "budi"
	limit := c.DefaultQuery("limit", "5") // "10", atau "5" kalau parameter limit gak dikirim
	c.JSON(200, gin.H{"nama": nama, "limit": limit})
}
```
`c.Query(...)` setara `r.URL.Query().Get(...)` di `net/http`, cuma nambah versi dengan default value (`c.DefaultQuery`).

## Baca Path Parameter

Contoh: `GET /user/5` (ambil user dengan ID 5)

```go
router.GET("/user/:id", getUserHandler)

func getUserHandler(c *gin.Context) {
	id := c.Param("id") // "5", tipe string
	c.JSON(200, gin.H{"id": id})
}
```
Path parameter di Gin ditandai `:nama` (titik dua), beda dari `net/http` Go 1.22+ yang pakai `{nama}` (kurung kurawal) dan diambil lewat `r.PathValue(...)`.

## Response dengan Status Code Eksplisit

```go
func getProductHandler(c *gin.Context) {
	id := c.Param("id")
	produk, ada := cariProduk(id)

	if !ada {
		c.JSON(404, gin.H{"message": "produk tidak ditemukan"})
		return
	}

	c.JSON(200, produk)
}
```

## Poin Penting

- Sama seperti `net/http`, `GET` idealnya tidak mengubah data di server (idempotent).
- `c.Param(...)` dan `c.Query(...)` selalu balikin `string` — perlu `strconv.Atoi(...)` dkk kalau butuh tipe lain.
- Konsep dasarnya identik dengan versi `net/http`, cuma penamaan method & cara nulis path parameter yang beda (`:id` vs `{id}`).
