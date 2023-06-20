package main

import (
	"golang-design-patterns/builder"
	"golang-design-patterns/chain-of-responsibility"
	"golang-design-patterns/decorator"
	"golang-design-patterns/factory"
	abstract_factory "golang-design-patterns/factory/abstract-factory"
	"golang-design-patterns/flyweight"
	"golang-design-patterns/mediator"
	"golang-design-patterns/observer"
	"golang-design-patterns/proxy"
	"golang-design-patterns/singleton"
	"golang-design-patterns/strategy"
)

func main() {
	hangisiCalistirilsin := 10

	switch hangisiCalistirilsin {
	case 1:
		builder.OrnekCalistir()
	case 2:
		decorator.OrnekCalistir()
	case 3:
		factory.OrnekCalistir()
	case 4:
		abstract_factory.OrnekCalistir()
	case 5:
		flyweight.OrnekCalistir()
	case 6:
		mediator.OrnekCalistir()
	case 7:
		observer.OrnekCalistir()
	case 8:
		chain_of_responsibility.OrnekCalistir()
	case 9:
		proxy.OrnekCalistir()
	case 10:
		singleton.OrnekCalistir()
	case 11:
		strategy.OrnekCalistir()
	}

}
