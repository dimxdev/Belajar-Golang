# 07 - For & Switch

Contoh praktik: [`src/07-for.go`](../src/07-for.go)

Go cuma punya **satu keyword** buat perulangan: `for` — nggak ada `while`/`do-while` terpisah kayak bahasa lain, semua bentuk perulangan pakai `for`.

## 1. Gaya Klasik (C-style)

```go
for i := 0; i < 7; i++ {
	fmt.Println(i)
}
```
Format: `init; kondisi; increment`.

## 2. Gaya Mirip `while`

```go
i := 0
for i < 7 {
	fmt.Println(i)
	i++
}
```
Cukup kondisi doang, tanpa init/increment eksplisit di baris `for`.

## 3. Infinite Loop

```go
for {
	// jalan terus sampai ada break/return
}
```

## 4. Loop dengan `range` (buat slice, map, string)

```go
for i, v := range someSlice {
	fmt.Println(i, v)
}
```
Dibahas lebih detail di materi [04 - Slice](04-slice.md) dan [05 - Map](05-map.md).

## `break` dan `continue`

```go
for i := 0; i < 10; i++ {
	if i == 3 {
		continue // skip iterasi ini, lanjut ke berikutnya
	}
	if i == 7 {
		break // langsung keluar dari loop
	}
	fmt.Println(i)
}
```

## Switch

```go
hari = strings.ToLower(hari)

switch hari {
case "senin":
	fmt.Println("ini hari senin")
case "selasa":
	fmt.Println("ini hari selasa")
default:
	fmt.Println(hari, "gadiajak wlee")
}
```
- Tiap `case` di Go **otomatis `break`** (beda dari C/JS yang butuh `break` manual biar nggak "jatuh" ke case berikutnya).
- `default` dijalanin kalau nggak ada `case` yang cocok.

**Switch tanpa kondisi** — jadi alternatif dari rangkaian `if-else if` yang panjang:
```go
switch {
case umur < 13:
	fmt.Println("anak-anak")
case umur < 18:
	fmt.Println("remaja")
default:
	fmt.Println("dewasa")
}
```

**`fallthrough`** — kalau butuh perilaku "jatuh ke case berikutnya" ala C/JS (jarang dipakai):
```go
switch angka {
case 1:
	fmt.Println("satu")
	fallthrough
case 2:
	fmt.Println("dua") // ikut ke-print meski angka == 1
}
```

## Poin Penting

- `for` itu satu-satunya keyword perulangan di Go, bentuknya fleksibel (klasik, while-style, infinite, range).
- Switch case Go otomatis `break`, kebalikan dari kebiasaan bahasa lain — kalau butuh perilaku "jatuh ke bawah", harus eksplisit pakai `fallthrough`.
