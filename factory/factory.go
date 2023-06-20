package factory

type Gardirop struct{}

func (g Gardirop) UstGetir(mevsim string) UstGiyilebilir {
	switch mevsim {
	case "Yaz":
		return &Tshirt{}
	case "Kış":
		return &Hirka{}
	}

	return nil
}
