package chain_of_responsibility

import "fmt"

func OrnekCalistir() {
	var (
		hamurMakinasi   HamurMakinasi
		sosMakinasi     SosMakinasi
		malzemeMakinasi MalzemeMakinasi
		pisirmeMakinasi PisirmeMakinasi

		sucukluKasarliPizza, karisikPizza Pizza
	)

	hamurMakinasi.SonrakiMakinayiBelirle(&sosMakinasi)
	sosMakinasi.SonrakiMakinayiBelirle(&malzemeMakinasi)
	malzemeMakinasi.SonrakiMakinayiBelirle(&pisirmeMakinasi)

	sucukluKasarliPizza = Pizza{
		hamur:      "İnce",
		sos:        "Normal",
		malzemeler: []PizzaMalzemeleri{Kasar, Sucuk},
	}
	hamurMakinasi.GoreviniYap(sucukluKasarliPizza)

	fmt.Println("------------------------")

	karisikPizza = Pizza{
		hamur:      "Normal",
		sos:        "Bol",
		malzemeler: []PizzaMalzemeleri{Kasar, Sucuk, Sosis, Zeytin, Misir, Biber},
	}
	hamurMakinasi.GoreviniYap(karisikPizza)

}
