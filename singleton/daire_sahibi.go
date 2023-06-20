package singleton

import (
	"sync"
)

type DaireSahibi struct {
	imza string
}

func (ds DaireSahibi) Imzala(wait *sync.WaitGroup) {
	defer wait.Done()

	dilekce = DilekceyiGetir()
	dilekce.Imzala(ds.imza)
}
