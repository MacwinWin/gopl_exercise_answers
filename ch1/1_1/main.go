// @author : microfat
// @time   : 09/10/20 15:38:46
// @File   : main.go

// Modify the echo program to also print os.Args[0], the name of the command that invoked it.package main

package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	s = os.Args[0]
	fmt.Println(s)
}
