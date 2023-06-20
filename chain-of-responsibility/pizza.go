package chain_of_responsibility

type Pizza struct {
	hamur      string
	sos        string
	malzemeler []PizzaMalzemeleri
}

type PizzaMalzemeleri string

const (
	Kasar  PizzaMalzemeleri = "Kasar"
	Sucuk                   = "Sucuk"
	Sosis                   = "Sosis"
	Zeytin                  = "Zeytin"
	Misir                   = "Misir"
	Biber                   = "Biber"
)
