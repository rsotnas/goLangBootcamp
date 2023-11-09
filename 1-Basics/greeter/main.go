package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
