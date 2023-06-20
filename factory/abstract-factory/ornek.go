package abstract_factory

import (
	"fmt"

	"golang-design-patterns/factory/abstract-factory/alt"
	"golang-design-patterns/factory/abstract-factory/ayak"
	"golang-design-patterns/factory/abstract-factory/ust"
)

func OrnekCalistir() {
	var (
		gardiropMerkezi GardiropMerkezi
		gardirop        Gardirop

		ustGiysi ust.UstGiyilebilir
		altGiysi alt.AltGiyilebilir
		ayakkabi ayak.AyakGiyilebilir
	)

	gardirop = gardiropMerkezi.GardiropGetir("Yaz")
	ustGiysi = gardirop.UstGiysiGetir()
	altGiysi = gardirop.AltGiysiGetir()
	ayakkabi = gardirop.AyakkabiGetir()
	ustGiysi.UstuneGiy()
	altGiysi.AltinaGiy()
	ayakkabi.AyaginaGiy()

	fmt.Println("------------------------")

	gardirop = gardiropMerkezi.GardiropGetir("Kış")
	gardirop.UstGiysiGetir().UstuneGiy()
	gardirop.AltGiysiGetir().AltinaGiy()
	gardirop.AyakkabiGetir().AyaginaGiy()
}
