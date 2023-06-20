package builder

type YuksekKalitePC struct {
	islemci    string
	ekranKarti string
	ram        string
	depolama   string
	anakart    string
}

func (pc *YuksekKalitePC) islemciEkle() {
	pc.islemci = "Yüksek kalite İşlemci"
}

func (pc *YuksekKalitePC) ekranKartiEkle() {
	pc.ekranKarti = "Yüksek kalite Ekran Kartı"
}

func (pc *YuksekKalitePC) ramEkle() {
	pc.ram = "Yüksek kalite Ram"
}

func (pc *YuksekKalitePC) depolamaEkle() {
	pc.depolama = "Yüksek kalite Depolama"
}

func (pc *YuksekKalitePC) anakartEkle() {
	pc.anakart = "Yüksek kalite Anakart"
}

func (pc *YuksekKalitePC) parcalariBirlestir() Bilgisayar {
	return Bilgisayar{
		islemci:    pc.islemci,
		ekranKarti: pc.ekranKarti,
		ram:        pc.ram,
		depolama:   pc.depolama,
		anakart:    pc.anakart,
	}
}
