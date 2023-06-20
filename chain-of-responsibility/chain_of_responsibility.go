package chain_of_responsibility

type Makina interface {
	GoreviniYap(pizza Pizza)
	SonrakiMakinayiBelirle(makina Makina)
}
