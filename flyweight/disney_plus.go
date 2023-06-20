package flyweight

import "fmt"

type DisneyPlus struct{}

func (n *DisneyPlus) GirisYap() {
	fmt.Println("Disney+'a giriliyor")
}
