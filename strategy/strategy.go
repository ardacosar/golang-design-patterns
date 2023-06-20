package strategy

type Asker interface {
	Giyin()
	TehcizatHazirla()
	HareketEt()
	EtkisizHaleGetir()
}

func OperasyonTarziBelirle(gorevTarzi string) Asker {
	switch gorevTarzi {
	case "Normal":
		return Normal{}
	case "Gizli":
		return Gizli{}
	default:
		panic("Asker bu operayon için uygun değil")
	}
}
