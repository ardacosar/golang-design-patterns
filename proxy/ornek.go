package proxy

import "fmt"

func OrnekCalistir() {
	var MuzKralligi Krallik

	fmt.Println("------------ Suçlu Talebi --------------")
	MuzKralligi.KrallaGorusmeTalepEt("Suçlu Yasin")

	fmt.Println("------------ Halk Talebi --------------")
	MuzKralligi.KrallaGorusmeTalepEt("Halktan biri Mehmet")

	fmt.Println("------------ Lord Talebi --------------")
	MuzKralligi.KrallaGorusmeTalepEt("Lord Cenk")
}
