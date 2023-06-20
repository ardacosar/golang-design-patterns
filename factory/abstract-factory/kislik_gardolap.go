package abstract_factory

import (
	"golang-design-patterns/factory/abstract-factory/alt"
	"golang-design-patterns/factory/abstract-factory/ayak"
	"golang-design-patterns/factory/abstract-factory/ust"
)

type KislikGardirop struct{}

func (kg *KislikGardirop) UstGiysiGetir() ust.UstGiyilebilir {
	return ust.Hirka{}
}

func (kg *KislikGardirop) AltGiysiGetir() alt.AltGiyilebilir {
	return alt.Pantolon{}
}

func (kg *KislikGardirop) AyakkabiGetir() ayak.AyakGiyilebilir {
	return ayak.Bot{}
}
