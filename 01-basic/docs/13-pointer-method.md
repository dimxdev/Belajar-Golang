# 13 - Pointer di Method

Contoh praktik: [`src/13-pointer-method.go`](../src/13-pointer-method.go)

Gabungan dari konsep pointer ([10](10-pointer.md)–[12](12-pointer-function.md)) dengan method ([09](09-method.md)) — pakai **pointer receiver** supaya method bisa mengubah struct aslinya.

## Pointer Receiver

```go
type Mahasiswa struct {
	Name string
	age  int
}

func (m *Mahasiswa) Ultah() {
	m.age++ // Go otomatis dereference, sama dengan (*m).age++
}

func (m *Mahasiswa) Update(name string, age int) {
	m.Name = name
	m.age = age
}
```
`(m *Mahasiswa)` adalah pointer receiver — method ini "nempel" ke **alamat** `Mahasiswa`, bukan salinannya, sehingga perubahan field di dalam method benar-benar mengubah struct aslinya.

## Pemanggilan — Otomatis Jadi Pointer

```go
mahasiswa1 := Mahasiswa{Name: "Dimas", age: 20}

mahasiswa1.Ultah() // Go otomatis convert jadi (&mahasiswa1).Ultah()
fmt.Println(mahasiswa1) // age jadi 21
```
Meskipun method didefinisikan dengan pointer receiver, cara memanggilnya tetap `mahasiswa1.Ultah()` biasa — Go otomatis mengambil alamat `mahasiswa1` di belakang layar.

## Kapan WAJIB Pointer Receiver (Bukan Cuma "Boleh")

1. Method perlu mengubah field struct aslinya.
2. Struct-nya berukuran besar (banyak field) — pointer receiver menghindari copy berulang tiap method dipanggil.
3. Struct itu berisi field yang tidak boleh/tidak bisa disalin (misal `sync.Mutex`).

## Konsistensi Receiver dalam Satu Tipe

Best practice: kalau **salah satu** method suatu tipe butuh pointer receiver, sebaiknya **semua** method tipe itu pakai pointer receiver juga — supaya konsisten dan menghindari kebingungan soal apakah suatu pemanggilan method mengubah data asli atau tidak.

```go
type Akun struct {
	saldo int
}

func (a *Akun) Setor(jumlah int)  { a.saldo += jumlah }
func (a *Akun) Tarik(jumlah int)  { a.saldo -= jumlah }
func (a *Akun) Saldo() int        { return a.saldo } // tetap pointer meski cuma baca, demi konsistensi
```

## Poin Penting

- Aturan sama seperti materi 09: **value receiver** untuk method yang hanya membaca, **pointer receiver** untuk method yang mengubah data.
- `variabel.Method()` dan `(&variabel).Method()` **setara** untuk pointer receiver — Go menangani konversinya otomatis.
- Pola ini dipakai luas di kode nyata, misal method `Save()`, `Update()`, `Delete()` pada struct model yang merepresentasikan baris database.
