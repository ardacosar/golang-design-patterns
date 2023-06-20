package decorator

import "fmt"

func OrnekCalistir() {
	var icindekiSivi string

	var suSisesi SuSisesi
	suSisesi.Doldur("Kirli Su")

	icindekiSivi = suSisesi.Ic()
	fmt.Println(icindekiSivi, "içildi")

	fmt.Println("------------------------")

	temizlikFiltreli := TemizlemeFiltresi{
		sise: &suSisesi,
	}
	icindekiSivi = temizlikFiltreli.Ic()
	fmt.Println(icindekiSivi, "içildi")

	fmt.Println("------------------------")

	aromaVeTemizlikFiltreli := AromaFiltresi{
		sise: &temizlikFiltreli,
	}
	icindekiSivi = aromaVeTemizlikFiltreli.Ic()
	fmt.Println(icindekiSivi, "içildi")
}
