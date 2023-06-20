package factory

import "fmt"

func OrnekCalistir() {
	var (
		ust      UstGiyilebilir
		gardirop Gardirop
	)

	ust = gardirop.UstGetir("Yaz")
	ust.UstuneGiy()

	fmt.Println("------------------------")

	ust = gardirop.UstGetir("Kış")
	ust.UstuneGiy()
}
