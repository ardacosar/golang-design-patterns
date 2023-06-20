package observer

import "fmt"

type Lamba struct{}

func (l Lamba) Aktifles() {
	fmt.Println("Lamba açılıyor")
}

func (l Lamba) Kapan() {
	fmt.Println("Lamba kapanıyor")
}
