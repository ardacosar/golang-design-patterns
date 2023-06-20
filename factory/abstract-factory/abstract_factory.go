package abstract_factory

type GardiropMerkezi struct{}

func (gm GardiropMerkezi) GardiropGetir(mevsim string) Gardirop {
	switch mevsim {
	case "Yaz":
		return &YazlikGardirop{}
	case "Kış":
		return &KislikGardirop{}
	}

	return nil
}
