package main

import (
	"fmt"
	"os"
)

func main() {
	flags, vals := parseFlags(os.Args[1:])
	fmt.Println(flags)
	fmt.Println(vals)
}
