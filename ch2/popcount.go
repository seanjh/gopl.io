package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/seanjh/gopl.io/ch2/popcount"
)

func main() {
	for _, val := range os.Args[1:] {
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		fmt.Printf("PopCount of %d (%b): %d\n", num, num, popcount.PopCount(uint64(num)))
	}
}
