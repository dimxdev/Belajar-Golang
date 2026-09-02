# 02 - Function

Contoh praktik: [`src/02-function.go`](../src/02-function.go)

## Function Dasar

```go
func AddNumbers(a int, b int) int {
	return a + b
}
```
- Parameter ditulis `nama tipe`. Kalau beberapa parameter berurutan tipenya sama, boleh disingkat: `func AddNumbers(a, b int) int`.
- Return type ditulis setelah daftar parameter.

## Function Tanpa Return Value

```go
func CetakSalam(nama string) {
	fmt.Println("Halo,", nama)
}
```

## Return Value Ganda

Pola paling khas di Go: function bisa return **lebih dari satu nilai sekaligus**, paling umum dipakai buat pola `(hasil, error)`:
```go
func Pembagian(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa dibagi nol")
	}
	return a / b, nil
}
```

## Named Return Value

Return value bisa dikasih nama, dan otomatis di-`return` tanpa perlu nulis eksplisit:
```go
func BagiNamed(a, b int) (hasil int, err error) {
	if b == 0 {
		err = fmt.Errorf("tidak bisa dibagi nol")
		return
	}
	hasil = a / b
	return
}
```

## Variadic Function (Jumlah Parameter Fleksibel)

```go
func Total(angka ...int) int {
	sum := 0
	for _, n := range angka {
		sum += n
	}
	return sum
}

Total(1, 2, 3)       // bisa dipanggil dengan jumlah argumen berapa aja
Total(1, 2, 3, 4, 5)
```

## Function sebagai Value / Parameter (First-Class Function)

Di Go, function bisa disimpan ke variabel, dikirim sebagai parameter, atau di-return dari function lain:
```go
func Operasi(a, b int, f func(int, int) int) int {
	return f(a, b)
}

hasil := Operasi(5, 3, AddNumbers)
```

## Anonymous Function & Closure

```go
tambah := func(a, b int) int {
	return a + b
}
fmt.Println(tambah(2, 3))
```

## Poin Penting

- Return ganda `(value, error)` adalah pola yang **paling sering** dipakai di Go, karena Go nggak punya exception — semua kemungkinan gagal harus dinyatakan lewat return value.
- Variadic (`...int`) berguna kalau jumlah argumen nggak pasti dari awal.
- Function first-class berguna buat pola callback, middleware (nanti relevan pas belajar HTTP handler), atau strategi pattern.
