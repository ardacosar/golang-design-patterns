package chain_of_responsibility

import "fmt"

type SosMakinasi struct {
	sonrakiMakina Makina
}

func (sm *SosMakinasi) GoreviniYap(pizza Pizza) {
	fmt.Println("Sos istege göre dökülüyor:", pizza.sos)
	//----------İşlemler----------
	fmt.Println("Sos başarıyla döküldü")
	if sm.sonrakiMakina != nil {
		sm.sonrakiMakina.GoreviniYap(pizza)
	}
}

func (sm *SosMakinasi) SonrakiMakinayiBelirle(makina Makina) {
	sm.sonrakiMakina = makina
}
