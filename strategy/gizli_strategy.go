package strategy

import "fmt"

type Gizli struct{}

func (g Gizli) Giyin() {
	fmt.Println("Hafif bir çelik yelek ve ortamda kamufle olunabilecek şekilde giyiliyor")
}
func (g Gizli) TehcizatHazirla() {
	fmt.Println("USP tactical ve şarjör hazırlanıyor. Susturucu takılıyor. ")
}
func (g Gizli) HareketEt() {
	fmt.Println("Ses çıkarmadan yavaşça harekete geçiliyor")
}
func (g Gizli) EtkisizHaleGetir() {
	fmt.Println("Düşman askeri sessizce etkisiz hale getirildi")
}
