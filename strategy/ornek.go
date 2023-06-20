package strategy

import "fmt"

func OrnekCalistir() {
	var Atilla Asker

	fmt.Println("------------ Normal Görev --------------")

	Atilla = OperasyonTarziBelirle("Normal")
	Atilla.Giyin()
	Atilla.TehcizatHazirla()
	Atilla.HareketEt()
	Atilla.EtkisizHaleGetir()

	fmt.Println("------------ Gizli Görev --------------")

	Atilla = OperasyonTarziBelirle("Gizli")
	Atilla.Giyin()
	Atilla.TehcizatHazirla()
	Atilla.HareketEt()
	Atilla.EtkisizHaleGetir()
}
