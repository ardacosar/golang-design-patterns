package proxy

import (
	"fmt"
	"strings"
)

type Vekil struct {
	kral Kral
}

func (v Vekil) Dinle(gorusmeci string) {

	if strings.Contains(gorusmeci, "Suçlu") {
		fmt.Println("Krala ulaşmana izin veremem, ona zarar vermeye çalışabilirsin")
		return
	}

	if strings.Contains(gorusmeci, "Halktan biri") {
		fmt.Println("Talebinizi not ettik, kral müsait olduğunda size haber verilecektir")
		return
	}

	if strings.Contains(gorusmeci, "Lord") {
		fmt.Println("Hoş geldiniz", gorusmeci)

		v.kral.Dinle(gorusmeci)
	}
}
