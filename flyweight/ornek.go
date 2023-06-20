package flyweight

import (
	"fmt"
)

func OrnekCalistir() {
	var hesap Hesap

	Ali := &Insan{
		hesaplar: make(map[string]Hesap),
	}
	fmt.Printf("Ali'nin Hesapları: %v\n", Ali.hesaplar)

	for i := 0; i < 3; i++ {
		hesap = Ali.HesabiniRicaEt("Netflix")
		hesap.GirisYap()
	}
	fmt.Printf("Ali'nin Hesapları: %v\n", Ali.hesaplar)

	for i := 0; i < 4; i++ {
		hesap = Ali.HesabiniRicaEt("Disney+")
		hesap.GirisYap()
	}
	fmt.Printf("Ali'nin Hesapları: %v\n", Ali.hesaplar)

	for i := 0; i < 5; i++ {
		hesap = Ali.HesabiniRicaEt("BluTV")
		hesap.GirisYap()
	}
	fmt.Printf("Ali'nin Hesapları: %v", Ali.hesaplar)
}
