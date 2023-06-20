package flyweight

import "fmt"

type Netflix struct{}

func (n *Netflix) GirisYap() {
	fmt.Println("Netflix'e giriliyor")
}
