package flyweight

import "fmt"

type BluTV struct{}

func (n *BluTV) GirisYap() {
	fmt.Println("BluTV'ye giriliyor")
}
