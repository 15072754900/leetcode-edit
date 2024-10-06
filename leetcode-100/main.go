package main

import (
	"fmt"
	"hufeng/demo/dumpstring"
)

func main() {
	dumpstring.DoSomething()
	a := dumpstring.DoAgain("")
	fmt.Println(a)
}
