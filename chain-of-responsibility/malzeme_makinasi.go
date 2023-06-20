package chain_of_responsibility

import "fmt"

type MalzemeMakinasi struct {
	sonrakiMakina Makina
}

func (mm *MalzemeMakinasi) GoreviniYap(pizza Pizza) {
	fmt.Println("İstenen malzemeler pizzaya ekleniyor:", pizza.malzemeler)
	//----------İşlemler----------
	fmt.Println("İstenen malzemeler pizzaya eklendi")
	if mm.sonrakiMakina != nil {
		mm.sonrakiMakina.GoreviniYap(pizza)
	}
}

func (mm *MalzemeMakinasi) SonrakiMakinayiBelirle(makina Makina) {
	mm.sonrakiMakina = makina
}
