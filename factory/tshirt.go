package factory

import "fmt"

type Tshirt struct{}

func (t Tshirt) UstuneGiy() {
	fmt.Println("Tshirt giyildi.")
}
