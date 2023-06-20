package abstract_factory

import (
	"golang-design-patterns/factory/abstract-factory/alt"
	"golang-design-patterns/factory/abstract-factory/ayak"
	"golang-design-patterns/factory/abstract-factory/ust"
)

type Gardirop interface {
	UstGiysiGetir() ust.UstGiyilebilir
	AltGiysiGetir() alt.AltGiyilebilir
	AyakkabiGetir() ayak.AyakGiyilebilir
}
