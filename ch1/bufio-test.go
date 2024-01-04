package main

import (
	"bufio"
	"fmt"
)

func main () {
	_, tokens, err := bufio.ScanWords([]byte("How now brown cow"), false)
	if err != nil {
		fmt.Printf("scan words failed: %v", err)
		return
	}

	for _, token := range tokens {
		fmt.Printf("token=%s\n", string(token))
	}
}
