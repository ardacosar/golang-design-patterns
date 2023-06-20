package mediator

import "fmt"

type Ogrenci struct {
	Ad           string
	OkulNumarasi string
	Hobi         string
	Pano         *OkulPanosu
}

func YeniOgrenci(ad, okulNo, hobi string, pano *OkulPanosu) Ogrenci {
	ogrenci := Ogrenci{
		Ad:           ad,
		OkulNumarasi: okulNo,
		Hobi:         hobi,
		Pano:         pano,
	}

	return ogrenci
}

func (o *Ogrenci) DuyuruYayimla(duyuru Duyuru) {
	fmt.Println(o.Ad, "bir duyuru yayımlıyor. Duyuru:", duyuru.Aciklama)
	o.Pano.PanoyaEkle(duyuru)
}

func (o *Ogrenci) DuyurularaGozat() {
	duyuru := o.Pano.UygunDuruyuGetir(*o)
	if duyuru != nil {
		fmt.Println(o.Ad, "kendine uygun duyuruyu buldu. Duyuru:", duyuru.Aciklama)
	} else {
		fmt.Println(o.Ad, "kendine uygun bir duyuru bulamadı")
	}
}
