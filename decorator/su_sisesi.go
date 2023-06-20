package decorator

import "fmt"

type SuSisesi struct {
	icindekiSu string
}

func (ss *SuSisesi) Ic() string {
	return ss.icindekiSu
}

func (ss *SuSisesi) Doldur(sivi string) {
	fmt.Println(sivi, "şişeye dolduruluyor...")
	ss.icindekiSu = sivi
}
