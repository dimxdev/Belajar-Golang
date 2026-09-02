# 09 - Method

Contoh praktik: [`src/09-method.go`](../src/09-method.go)

Method adalah function yang "nempel" ke suatu tipe (biasanya struct), ditulis dengan **receiver** di antara `func` dan nama method-nya. Go tidak punya class, tapi method + struct bisa memberikan efek serupa OOP dasar.

## Value Receiver (Read-Only)

```go
type Employee struct {
	Name string
	age  int
}

func (e Employee) Sapa() string {
	return "Halo, saya " + e.Name
}
```
`(e Employee)` adalah **value receiver** — `e` adalah **salinan/copy** dari struct aslinya. Cocok kalau method cuma **baca** data.

## Pointer Receiver (Bisa Mengubah Data Asli)

```go
func (e *Employee) UbahNama(namaBaru string) {
	e.Name = namaBaru
}
```
`(e *Employee)` adalah **pointer receiver** — `e` adalah pointer ke struct aslinya, jadi perubahan di dalam method beneran mengubah struct aslinya.

```go
employe1.UbahNama("Dimas H") // Go otomatis convert jadi (&employe1).UbahNama(...)
```

## Method pada Named Type Non-Struct

Method tidak cuma bisa nempel ke struct — tipe custom apapun bisa punya method:
```go
type Age int

func (a Age) isAdult() bool {
	return a >= 18
}
```

## Method dengan Parameter & Return Value

```go
func (e Employee) UbahRole(roleBaru string) Employee {
	e.role = roleBaru
	return e // return salinan yang sudah diubah (karena value receiver)
}
```

## Interface & Method (Sekilas)

Method itu yang bikin sebuah tipe "mengimplementasikan" sebuah interface — cukup punya method dengan signature yang cocok, tanpa keyword `implements` eksplisit:
```go
type Penyapa interface {
	Sapa() string
}

// Employee otomatis "menjadi" Penyapa karena sudah punya method Sapa() string
```

## Poin Penting: Kapan Pakai Value vs Pointer Receiver?

- **Value receiver** → method cuma **baca** data, tidak perlu mengubah struct aslinya.
- **Pointer receiver** → method perlu **mengubah** field struct aslinya, atau struct-nya besar (menghindari copy berulang tiap dipanggil).
- Best practice: kalau salah satu method di suatu tipe butuh pointer receiver, biasanya **semua** method di tipe itu dibuat pointer receiver juga, biar konsisten.
- `variabel.Method()` otomatis di-convert Go jadi `(&variabel).Method()` kalau method-nya pakai pointer receiver — tidak perlu ambil alamat manual saat memanggil.
