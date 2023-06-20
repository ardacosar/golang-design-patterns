package builder

type Personel struct {
	pcSetup PCSetup
}

func (d *Personel) SecilenKasa(b PCSetup) {
	d.pcSetup = b
}

func (d *Personel) KasayiTopla() Bilgisayar {
	d.pcSetup.islemciEkle()
	d.pcSetup.ekranKartiEkle()
	d.pcSetup.ramEkle()
	d.pcSetup.depolamaEkle()
	d.pcSetup.anakartEkle()
	return d.pcSetup.parcalariBirlestir()
}
