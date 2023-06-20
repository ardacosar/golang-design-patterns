package builder

type PCSetup interface {
	islemciEkle()
	ekranKartiEkle()
	ramEkle()
	depolamaEkle()
	anakartEkle()
	parcalariBirlestir() Bilgisayar
}

func PCSetupSecimi(istek string) PCSetup {
	switch istek {
	case "Oyun oynamak", "İş yapmak":
		return &YuksekKalitePC{}
	case "Hobi":
		return &OrtaKalitePC{}
	case "Film izlemek", "Dizi izlemek":
		return &DusukKalitePC{}
	}

	return nil
}
