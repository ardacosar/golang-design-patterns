package singleton

import (
	"fmt"
	"sync"
)

var dilekce *Dilekce

var apartmanYoneticisi sync.Once

func DilekceyiGetir() *Dilekce {
	if dilekce == nil {
		fmt.Println("Dilekçe boş, yaratılmak için harekete geçiliyor")
		apartmanYoneticisi.Do(
			func() {
				fmt.Println("Dilekçe ilk kez yazıldı")
				dilekce = &Dilekce{}
			},
		)
	} else {
		fmt.Println("Dilekçe zaten var")
	}

	return dilekce
}
