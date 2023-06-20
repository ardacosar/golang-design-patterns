package builder

type OrtaKalitePC struct {
	islemci    string
	ekranKarti string
	ram        string
	depolama   string
	anakart    string
}

func (pc *OrtaKalitePC) islemciEkle() {
	pc.islemci = "Orta kalite İşlemci"
}

func (pc *OrtaKalitePC) ekranKartiEkle() {
	pc.ekranKarti = "Orta kalite Ekran Kartı"
}

func (pc *OrtaKalitePC) ramEkle() {
	pc.ram = "Orta kalite Ram"
}

func (pc *OrtaKalitePC) depolamaEkle() {
	pc.depolama = "Orta kalite Depolama"
}

func (pc *OrtaKalitePC) anakartEkle() {
	pc.anakart = "Orta kalite Anakart"
}

func (pc *OrtaKalitePC) parcalariBirlestir() Bilgisayar {
	return Bilgisayar{
		islemci:    pc.islemci,
		ekranKarti: pc.ekranKarti,
		ram:        pc.ram,
		depolama:   pc.depolama,
		anakart:    pc.anakart,
	}
}
