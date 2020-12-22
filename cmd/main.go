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
	var min int
	var max int
	var code string

	flag.StringVar(&mode, "mode", "numbers", "mode for code generation")
	flag.IntVar(&length, "length", 6, "length of generated code")
	flag.IntVar(&min, "min", 1, "minimal number")
	flag.IntVar(&max, "max", 10, "maximal number")
	flag.Parse()

	switch mode {
	case "random":
		code = onetimecode.RandomNumber(min, max)
	case "numbers":
		code = onetimecode.NumberCode(length)
	case "alphanum":
		code = onetimecode.AlphaNumberCode(length)
	case "alphanumuc":
		code = onetimecode.AlphaNumberUcCode(length)
	case "uuid":
		code = onetimecode.Uuid()
	}

	fmt.Println(code)
}
