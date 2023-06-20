package chain_of_responsibility

import "fmt"

type HamurMakinasi struct {
	sonrakiMakina Makina
}

func (hm *HamurMakinasi) GoreviniYap(pizza Pizza) {
	fmt.Println("Hamur kalınlık tercihine göre açılıyor:", pizza.hamur)
	//----------İşlemler----------
	fmt.Println("Hamur başarıyla açıldı")
	if hm.sonrakiMakina != nil {
		hm.sonrakiMakina.GoreviniYap(pizza)
	}
}

func (hm *HamurMakinasi) SonrakiMakinayiBelirle(makina Makina) {
	hm.sonrakiMakina = makina
}
