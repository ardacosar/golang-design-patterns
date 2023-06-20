package builder

import "fmt"

func OrnekCalistir() {
	var (
		PC           Bilgisayar
		topSakalReis Personel
	)

	fmt.Println("Oyun oynamak için PC tercihi yapılıyor...")
	oyunPc := PCSetupSecimi("Oyun oynamak")

	topSakalReis.SecilenKasa(oyunPc)
	PC = topSakalReis.KasayiTopla()
	fmt.Printf("İşlemci: %s\n", PC.islemci)
	fmt.Printf("Ekran Kartı: %s\n", PC.ekranKarti)
	fmt.Printf("Ram: %s\n", PC.ram)
	fmt.Printf("Depolama: %s\n", PC.depolama)
	fmt.Printf("Anakart: %s\n", PC.anakart)

	fmt.Println("------------------------")

	fmt.Println("Dizi izlemek için PC tercihi yapılıyor...")
	diziPc := PCSetupSecimi("Dizi izlemek")

	topSakalReis.SecilenKasa(diziPc)
	PC = topSakalReis.KasayiTopla()
	fmt.Printf("İşlemci: %s\n", PC.islemci)
	fmt.Printf("Ekran Kartı: %s\n", PC.ekranKarti)
	fmt.Printf("Ram: %s\n", PC.ram)
	fmt.Printf("Depolama: %s\n", PC.depolama)
	fmt.Printf("Anakart: %s\n", PC.anakart)
}
