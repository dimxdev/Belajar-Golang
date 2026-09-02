# 06 - If Else

Contoh praktik: [`src/06-if-else.go`](../src/06-if-else.go)

## If Sederhana

```go
if umur >= 17 {
	fmt.Println("udah dewasa")
}
```
Kondisi **tidak pakai kurung** `()` (beda dari banyak bahasa lain), tapi kurung kurawal `{}` **wajib**, meski body-nya cuma 1 baris.

## If - Else

```go
if umur >= 17 {
	fmt.Println("udah dewasa")
} else {
	fmt.Println("masi anak anak")
}
```

## If - Else If - Else

```go
if umur >= 17 {
	fmt.Println("udah dewasa")
} else if umur >= 10 {
	fmt.Println("lumayan lah")
} else {
	fmt.Println("masi bocil")
}
```

## Deklarasi Variabel Langsung di If (Init Statement)

```go
if nilai := 99; nilai <= 80 {
	fmt.Println("busyukkk")
} else {
	fmt.Println("bagus2")
}
```
Variabel `nilai` dideklarasi langsung di statement if, dipisah `;` dari kondisinya. Scope-nya cuma berlaku di dalam blok `if`/`else if`/`else` itu, tidak bisa diakses di luar.

## Operator Logika

| Operator | Arti |
|---|---|
| `&&` | AND — kedua kondisi harus benar |
| `\|\|` | OR — salah satu kondisi cukup benar |
| `!` | NOT — membalik nilai boolean |

```go
if umur >= 17 && umur < 60 {
	fmt.Println("usia produktif")
}
```

## Operator Perbandingan

| Operator | Arti |
|---|---|
| `==` | sama dengan |
| `!=` | tidak sama dengan |
| `>`, `<`, `>=`, `<=` | lebih besar/kecil (sama dengan) |

## Poin Penting

- Kurung kurawal `{}` wajib, meski body-nya cuma 1 statement — beda dari C/JS yang mengizinkan 1 statement tanpa `{}`.
- Pola "init + kondisi" di if (`if x := ...; kondisi`) sangat sering dipakai bareng function yang return `(value, error)`, misal `if err := doSomething(); err != nil { ... }`.
