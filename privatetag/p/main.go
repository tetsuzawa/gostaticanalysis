package main

import (
	"flag"
	"fmt"
)

func main() {
	numArgs := flag.Int("numvars", 5, "number of args which is allowed to pass to func")
	fmt.Println(numArgs)
	fmt.Println(flag.Parse)

}
