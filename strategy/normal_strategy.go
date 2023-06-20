package strategy

import "fmt"

type Normal struct{}

func (n Normal) Giyin() {
	fmt.Println("Sağlam bir çelik yelek ve havalı asker kostümü giyiliyor")
}
func (n Normal) TehcizatHazirla() {
	fmt.Println("M4A1 ve şarjörleri hazırlanıyor")
}
func (n Normal) HareketEt() {
	fmt.Println("Hedef konuma doğru harekete geçiliyor")
}
func (n Normal) EtkisizHaleGetir() {
	fmt.Println("Düşman askeri ortadan kaldırıldı")
}
