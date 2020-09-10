// @author : microfat
// @time   : 09/10/20 16:11:18
// @File   : main.go

// Modify the echo program to print the index and value of each of its arguments, one per line.package main

package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Println(index, value)
	}
}
