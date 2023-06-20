package observer

import "fmt"

func OrnekCalistir() {
	var (
		akilliEv AkilliEv

		lamba1      Lamba
		televizyon1 Televizyon
		klima1      Klima
	)

	akilliEv.SistemeBagla(lamba1)
	akilliEv.SistemeBagla(televizyon1)
	akilliEv.SistemeBagla(klima1)

	akilliEv.HerseyiAc()
	fmt.Println("------------------------")
	akilliEv.HerseyiKapat()

	fmt.Println("================================")

	akilliEv.SistemdenCikar(klima1)

	akilliEv.HerseyiAc()
	fmt.Println("------------------------")
	akilliEv.HerseyiKapat()
}
