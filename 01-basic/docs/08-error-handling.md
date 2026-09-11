# 08 - Error Handling

Contoh praktik: [`src/08-error-handling.go`](../src/08-error-handling.go)

Go **tidak punya** try/catch/exception seperti kebanyakan bahasa lain. Pendekatannya: function yang bisa gagal, return `error` sebagai nilai biasa, dan pemanggil **wajib cek manual**.

## Bikin Error

**`errors.New`** — buat pesan error statis:
```go
import "errors"

func cekStok(stok int) error {
	if stok <= 0 {
		return errors.New("stok habis")
	}
	return nil
}
```

**`fmt.Errorf`** — buat pesan error dinamis (bisa sisip variabel, mirip `Sprintf`):
```go
func Bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("%d tidak bisa dibagi dengan 0", a)
	}
	return a / b, nil
}
```

## Pola Cek Error (Fail Fast)

```go
hasil, err := Bagi(10, 0)
if err != nil {
	fmt.Println("Error:", err)
	return // langsung berhenti, jangan lanjut kalau ada error
}
fmt.Println("Hasil:", hasil)
```

## Error di Dalam Loop

```go
for _, stok := range stoks {
	err := cekStok(stok)
	if err != nil {
		fmt.Println("Stok", stok, "->", err)
		continue // skip iterasi ini doang, bukan berhenti total
	}
	fmt.Println("Stok", stok, "-> Aman")
}
```

## Error Wrapping (`%w`)

Buat "membungkus" error asli sambil nambahin konteks, tanpa kehilangan info error aslinya:
```go
if err != nil {
	return fmt.Errorf("gagal ambil data user: %w", err)
}
```
Error yang di-wrap gitu bisa dicek/dibongkar lagi pakai `errors.Is` dan `errors.As`.

## Custom Error Type

Kalau butuh error yang bawa data tambahan (bukan cuma pesan teks):
```go
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}
```
Tipe apapun yang punya method `Error() string` otomatis "menjadi" tipe `error` (ini konsep interface, dibahas lebih lanjut nanti).

## `panic` dan `recover` (Jarang Dipakai)

Beda dari `error` biasa (yang expected/predictable), `panic` dipakai buat kondisi yang **beneran fatal/tidak terduga**:
```go
func mayPanic() {
	panic("sesuatu yang fatal terjadi")
}

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	mayPanic()
}
```
Di Go, `panic`/`recover` **bukan** pengganti error handling biasa — pola `if err != nil` tetap yang utama, `panic` cuma buat kasus yang benar-benar exceptional (bug, out-of-bounds, dll).

## Poin Penting

- Pola `if err != nil { ...; return }` akan **terus muncul** di hampir semua kode Go, termasuk nanti di handler REST API.
- `return` dipakai kalau error bikin proses harus **berhenti total**; `continue` kalau cuma iterasi itu doang yang perlu di-skip.
- Selalu cek `err != nil` **sebelum** pakai nilai hasil lainnya — nilai hasil biasanya tidak valid/zero value kalau error terjadi.
