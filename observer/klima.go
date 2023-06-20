package observer

import "fmt"

type Klima struct{}

func (k Klima) Aktifles() {
	fmt.Println("Klima açılıyor")
}

func (k Klima) Kapan() {
	fmt.Println("Klima kapanıyor")
}
