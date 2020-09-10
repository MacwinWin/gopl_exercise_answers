// @author : microfat
// @time   : 09/10/20 16:22:51
// @File   : main.go

// Experiment to measure the difference in running time between our
// potentially inefficient versions and the one that uses strings.Join.
// (Section 1.6 illustrates part of the time package, and Section 11.4
//	shows how to write benchmark tests for systematic performance evaluation.)

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// slow
	startSlow := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.9fs elapsed\n", time.Since(startSlow).Seconds())
	// fast
	startFast := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.9fs elapsed\n", time.Since(startFast).Seconds())
}
