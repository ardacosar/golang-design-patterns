package decorator

import "fmt"

type AromaFiltresi struct {
	sise Sise
}

func (af *AromaFiltresi) Ic() string {
	fmt.Println("Aroma filtresi takıldı")

	yeniSivi := "Aromalı " + af.sise.Ic()
	return yeniSivi
}
