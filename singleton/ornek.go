package singleton

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func OrnekCalistir() {
	var apartmanSakinleri []DaireSahibi
	kisiSayisi := 5

	for i := 1; i <= kisiSayisi; i++ {
		apartmanSakinleri = append(apartmanSakinleri,
			DaireSahibi{
				imza: fmt.Sprintf("imza-%d", i),
			})
	}
	wait.Add(kisiSayisi)

	for _, daireSahibi := range apartmanSakinleri {
		go daireSahibi.Imzala(&wait)
	}

	wait.Wait()

	fmt.Println("İmzalar:", dilekce.imzalar)
	fmt.Println("İmza sayısı:", len(dilekce.imzalar))
}
