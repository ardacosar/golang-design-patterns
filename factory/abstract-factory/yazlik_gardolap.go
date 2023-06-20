package abstract_factory

import (
	"golang-design-patterns/factory/abstract-factory/alt"
	"golang-design-patterns/factory/abstract-factory/ayak"
	"golang-design-patterns/factory/abstract-factory/ust"
)

type YazlikGardirop struct{}

func (yg *YazlikGardirop) UstGiysiGetir() ust.UstGiyilebilir {
	return ust.Tshirt{}
}

func (yg *YazlikGardirop) AltGiysiGetir() alt.AltGiyilebilir {
	return alt.Sort{}
}

func (yg *YazlikGardirop) AyakkabiGetir() ayak.AyakGiyilebilir {
	return ayak.Sandalet{}
}
