package proxy

import "fmt"

type Krallik struct{}

func (k Krallik) KrallaGorusmeTalepEt(gorusmeci string) {
	var dinleyici Dinleyici

	//  dinleyici = Kral{Isim: "Süleyman"}
	dinleyici = Vekil{
		kral: Kral{Isim: "Süleyman"},
	}

	fmt.Println("Kapıya ulaşıldı...")
	dinleyici.Dinle(gorusmeci)
}
