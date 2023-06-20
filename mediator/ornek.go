package mediator

func OrnekCalistir() {
	var pano OkulPanosu
	var mahmut, turgut, cenk, cansu Ogrenci

	turgut = YeniOgrenci("Turgut", "967", "Tiyatro", &pano)
	cenk = YeniOgrenci("Cenk", "412", "Futbol", &pano)
	cansu = YeniOgrenci("Cansu", "119", "Gitar", &pano)
	mahmut = YeniOgrenci("Mahmut", "235", "Tiyatro", &pano)

	mahmut.DuyurularaGozat()
	cenk.DuyurularaGozat()

	turgut.DuyuruYayimla(Duyuru{
		Aciklama:   "Tiyatro kulübü açıyoruz, katılımcıları bekleriz",
		IletisimNo: "0222 222 22 22",
	})

	cansu.DuyurularaGozat()
	mahmut.DuyurularaGozat()
}
