package observer

import (
	"fmt"
	"reflect"
)

type AkilliEv struct {
	urunler map[string]AkilliUrun
}

func (sk *AkilliEv) SistemeBagla(urun AkilliUrun) {
	if sk.urunler == nil {
		sk.urunler = make(map[string]AkilliUrun, 0)
	}

	urununAdi := reflect.TypeOf(urun).Name()
	sk.urunler[urununAdi] = urun
	fmt.Println("Sisteme baglandi:", urununAdi)
}

func (sk *AkilliEv) SistemdenCikar(urun AkilliUrun) {
	urununAdi := reflect.TypeOf(urun).Name()
	delete(sk.urunler, urununAdi)
	fmt.Println("Sistemle baglantisi kesildi:", urununAdi)
}

func (sk *AkilliEv) HerseyiAc() {
	fmt.Println("Sistem aktifleşiyor...")
	for _, urun := range sk.urunler {
		urun.Aktifles()
	}
}

func (sk *AkilliEv) HerseyiKapat() {
	fmt.Println("Sistem kapatılıyor...")
	for _, urun := range sk.urunler {
		urun.Kapan()
	}
}
