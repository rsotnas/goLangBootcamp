package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("secret/file.txt")
	fmt.Println("Dir: ", dir)
	fmt.Println("File: ", file)
}
