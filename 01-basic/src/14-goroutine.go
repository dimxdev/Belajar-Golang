package src

import "fmt"

func Sapa(selesai chan bool) {
	fmt.Println("Hello iam from go")
	selesai <- true
}

func MainSapa() {
	selesaii:= make(chan bool)

	go Sapa(selesaii)
	<- selesaii

	fmt.Println("selesai dari main")
	fmt.Println()
}


func MasakTelur(masak1 chan bool) {
	fmt.Println("masak telur selesai")
	masak1 <- true
}

func MasakMie(masak2 chan bool) {
	fmt.Println("masak Mie selesai")
	masak2 <- true
}

func MainMasak(){
	masak:= make(chan bool)

	go MasakMie(masak)
	go MasakTelur(masak)

	<- masak
	<- masak

	fmt.Println("masak selesai")
	fmt.Println()
}


func HitungKuadrat(hasil chan int, angka int) {
	hasil <- angka * angka
}

func MainHitungKuadrat() {
	hasil := make(chan int)
	angka := 20

	go HitungKuadrat(hasil, angka)
	nilai := <- hasil
	fmt.Printf("nilai kuadrat dari %d adalah %d \n" , angka, nilai) 
	fmt.Println("hitung selesai")

	hasil2 := make(chan int)
	angka2 := []int{2, 30, 50, 30}

	for _, a := range angka2 {
		go HitungKuadrat(hasil2, a)
	}

	for i := 0; i < len(angka2); i++ {
		fmt.Println(<-hasil2)
	}
}

// latihan
func KirimPesan(pesan chan string) {
	pesan <- "hello dari goroutine"
}

func MainKirimPesan() {
	pesan := make(chan string)
	go KirimPesan(pesan)

	fmt.Println(<-pesan)
}


