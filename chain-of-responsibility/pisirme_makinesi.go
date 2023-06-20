package chain_of_responsibility

import "fmt"

type PisirmeMakinasi struct {
	sonrakiMakina Makina
}

func (pm *PisirmeMakinasi) GoreviniYap(pizza Pizza) {
	fmt.Println("Pişirme işlemi başladı")
	//----------İşlemler----------
	fmt.Println("Pişirme işlemi tamamlandı")
	if pm.sonrakiMakina != nil {
		pm.sonrakiMakina.GoreviniYap(pizza)
	}
}

func (pm *PisirmeMakinasi) SonrakiMakinayiBelirle(makina Makina) {
	pm.sonrakiMakina = makina
}
