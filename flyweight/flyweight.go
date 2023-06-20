package flyweight

type Insan struct {
	hesaplar map[string]Hesap
}

func (i *Insan) HesabiniRicaEt(platform string) Hesap {
	if i.hesaplar[platform] != nil {
		return i.hesaplar[platform]
	}

	switch platform {
	case "Netflix":
		i.hesaplar[platform] = &Netflix{}
		return i.hesaplar[platform]
	case "Disney+":
		i.hesaplar[platform] = &DisneyPlus{}
		return i.hesaplar[platform]
	case "BluTV":
		i.hesaplar[platform] = &BluTV{}
		return i.hesaplar[platform]
	}

	return nil
}
