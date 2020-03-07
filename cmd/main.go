package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/jakoubek/onetimecode"
)

var source = rand.NewSource(time.Now().UnixNano())

func main() {
	var mode string
	var length int
	var code string

	flag.StringVar(&mode, "mode", "numbers", "mode for code generation")
	flag.IntVar(&length, "length", 6, "length of generated code")
	flag.Parse()

	switch mode {
	case "numbers":
		code = onetimecode.NumberCode(length)
	case "alphanum":
		code = onetimecode.AlphaNumberCode(length)
	case "alphanumuc":
		code = onetimecode.AlphaNumberUcCode(length)
	}

	fmt.Println(code)
}
