package mediator

import (
	"fmt"
	"strings"
)

type OkulPanosu struct {
	duyurular []Duyuru
}

func (op *OkulPanosu) PanoyaEkle(duyuru Duyuru) {
	op.duyurular = append(op.duyurular, duyuru)
	fmt.Println("Duyuru panoya asıldı.")
}

func (op *OkulPanosu) UygunDuruyuGetir(ogrenci Ogrenci) *Duyuru {
	for _, duyuru := range op.duyurular {
		if strings.Contains(duyuru.Aciklama, ogrenci.Hobi) {
			return &duyuru
		}
	}
	return nil
}
