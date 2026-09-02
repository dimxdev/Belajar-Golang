# 14 - Goroutine & Channel

Contoh praktik: [`src/14-goroutine.go`](../src/14-goroutine.go)

Goroutine adalah cara Go menjalankan function secara **concurrent** (bersamaan), cukup dengan menambahkan keyword `go` di depan pemanggilan function. Channel adalah "pipa" yang dipakai goroutine untuk berkomunikasi/bertukar data satu sama lain dengan aman.

## Goroutine Dasar

```go
func sapa() {
	fmt.Println("halo dari goroutine")
}

go sapa() // dijalankan di goroutine terpisah, TIDAK blocking
fmt.Println("lanjut duluan")
```
`go sapa()` langsung lanjut ke baris berikutnya tanpa menunggu `sapa` selesai — kalau function utama (`main`) keburu selesai/keluar, goroutine yang belum sempat jalan bisa saja tidak ter-eksekusi sama sekali.

## Channel sebagai Sinyal (Sinkronisasi)

```go
func Sapa(selesai chan bool) {
	fmt.Println("Hello dari goroutine")
	selesai <- true // kirim sinyal "sudah selesai"
}

selesai := make(chan bool)
go Sapa(selesai)
<-selesai // blocking, tunggu sampai ada sinyal masuk
fmt.Println("lanjut setelah goroutine selesai")
```
- `<-channel` (tanpa ditampung ke variabel) = blocking, menunggu sampai ada nilai masuk.
- Kalau ada **N** goroutine yang masing-masing mengirim 1 sinyal ke channel yang sama, pembacanya juga perlu membaca channel itu **N kali** supaya semuanya benar-benar ditunggu.

## Channel Membawa Data

```go
func hitungKuadrat(hasil chan int, angka int) {
	hasil <- angka * angka
}

hasil := make(chan int)
go hitungKuadrat(hasil, 5)
nilai := <-hasil // menunggu SEKALIGUS menangkap hasilnya
fmt.Println(nilai) // 25
```

## Pola: Launch Semua Dulu, Baru Kumpulkan Hasil

```go
angka := []int{2, 30, 50}
hasil := make(chan int)

for _, a := range angka {
	go hitungKuadrat(hasil, a) // launch SEMUA goroutine dulu (concurrent)
}
for i := 0; i < len(angka); i++ {
	fmt.Println(<-hasil) // baru kumpulkan hasilnya, di loop terpisah
}
```
- Loop pertama hanya **launching**, tidak menunggu — semua goroutine sempat berjalan bersamaan.
- Loop kedua baru **menunggu & membaca hasil**, sebanyak jumlah goroutine yang dilaunch.
- **Urutan hasil tidak dijamin** sama dengan urutan input asli — tergantung goroutine mana yang lebih dulu selesai (konsekuensi wajar dari concurrency).
- ⚠️ Kalau launch (`go ...`) dan tunggu (`<-channel`) ditaruh dalam **iterasi loop yang sama**, goroutine-nya berjalan **berurutan** (bukan bersamaan), karena tiap iterasi memblokir sebelum sempat melaunch goroutine berikutnya — efeknya jadi sama seperti tanpa `go` sama sekali.

## Buffered Channel

Channel biasa (unbuffered) memaksa pengirim menunggu sampai ada penerima. Buffered channel punya kapasitas, jadi pengirim tidak langsung blocking selama buffer belum penuh:
```go
ch := make(chan int, 3) // kapasitas 3
ch <- 1
ch <- 2
ch <- 3 // masih tidak blocking, buffer belum penuh
// ch <- 4 akan blocking karena buffer sudah penuh
```

## `sync.WaitGroup` — Alternatif Selain Channel Counting Manual

Untuk goroutine dalam jumlah banyak/dinamis, menghitung `<-channel` manual jadi tidak praktis. `sync.WaitGroup` adalah cara lebih umum untuk menunggu banyak goroutine:
```go
var wg sync.WaitGroup

for _, a := range angka {
	wg.Add(1)
	go func(n int) {
		defer wg.Done()
		fmt.Println(n * n)
	}(a)
}
wg.Wait() // tunggu semua goroutine selesai
```

## Poin Penting

- `go` di depan pemanggilan function = jalankan di goroutine baru, tidak blocking.
- Channel dipakai untuk dua hal: **sinkronisasi** (menunggu sinyal selesai) dan **transfer data** (mengirim hasil kerja goroutine).
- Membaca channel (`<-channel`) harus **sebanyak** jumlah goroutine yang mengirim ke situ — kalau tidak, goroutine yang kirimannya tak terbaca akan blocking selamanya (disebut *goroutine leak*).
- Kalau perlu tahu data berasal dari input yang mana (terutama jika ada kemungkinan duplikat), channel perlu membawa **pasangan data** (misal lewat struct), bukan cuma hasil akhirnya saja.
