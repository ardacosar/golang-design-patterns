package proxy

import "fmt"

type Kral struct {
	Isim string
}

func (k Kral) Dinle(gorusmeci string) {
	fmt.Println("İsteğini değerlendireceğim", gorusmeci)
}
