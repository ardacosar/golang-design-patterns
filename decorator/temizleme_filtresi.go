package decorator

import (
	"fmt"
	"strings"
)

type TemizlemeFiltresi struct {
	sise Sise
}

func (tf *TemizlemeFiltresi) Ic() string {
	fmt.Println("Temizlik filtresi takıldı")

	siseninIcindekiSivi := tf.sise.Ic()
	yeniSivi := strings.ReplaceAll(siseninIcindekiSivi, "Kirli ", "")
	return yeniSivi
}
