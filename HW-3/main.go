package main

import (
	//"HW-3/openfile"
	"HW-3/topwords"
	"fmt"
)

func main() {

	var p = fmt.Println

	/*s, err := openfile.OpenFile()

	if err != nil {
	 p(err)
	 return
	}*/
	//s :="camel rat leopard hamster giraffe lion rat ostrich panther rat lion hamster elephant cheetah kangaroo lion tiger"

	//rat 4 lion 3 giraffe 3  monkey 2

	s := `giraffe cat rat cheetah dog monkey elephant rat ferret lion
 giraffe pig hamster rat kangaroo kitten giraffe leopard lion monkey ostrich rat tiger lion turtle`

	// s :="111 222 111 222 111 333"
	//s :="cat and dog one dog two cats and one man"

	p(topwords.Top10(s))

}
