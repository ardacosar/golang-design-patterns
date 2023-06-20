package observer

import "fmt"

type Televizyon struct{}

func (t Televizyon) Aktifles() {
	fmt.Println("Televizyon açılıyor")
}

func (t Televizyon) Kapan() {
	fmt.Println("Televizyon kapanıyor")
}
