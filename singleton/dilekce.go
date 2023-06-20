package singleton

import (
	"fmt"
	"sync"
)

type Dilekce struct {
	imzalar []string
}

var gorevli sync.Mutex

func (d *Dilekce) Imzala(imza string) {
	gorevli.Lock()
	defer gorevli.Unlock()

	d.imzalar = append(d.imzalar, imza)
	fmt.Println("Dilekce imzalandi:", imza)
}
