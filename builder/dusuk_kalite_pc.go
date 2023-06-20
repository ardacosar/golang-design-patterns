package builder

type DusukKalitePC struct {
	islemci    string
	ekranKarti string
	ram        string
	depolama   string
	anakart    string
}

func (pc *DusukKalitePC) islemciEkle() {
	pc.islemci = "Düşük kalite İşlemci"
}

func (pc *DusukKalitePC) ekranKartiEkle() {
	pc.ekranKarti = "Düşük kalite Ekran Kartı"
}

func (pc *DusukKalitePC) ramEkle() {
	pc.ram = "Düşük kalite Ram"
}

func (pc *DusukKalitePC) depolamaEkle() {
	pc.depolama = "Düşük kalite Depolama"
}

func (pc *DusukKalitePC) anakartEkle() {
	pc.anakart = "Düşük kalite Anakart"
}

func (pc *DusukKalitePC) parcalariBirlestir() Bilgisayar {
	return Bilgisayar{
		islemci:    pc.islemci,
		ekranKarti: pc.ekranKarti,
		ram:        pc.ram,
		depolama:   pc.depolama,
		anakart:    pc.anakart,
	}
}
