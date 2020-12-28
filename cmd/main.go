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
	flag.IntVar(&length, "length", -1, "length of generated code")
	flag.IntVar(&min, "min", -1, "minimal number")
	flag.IntVar(&max, "max", -1, "maximal number")
	flag.Parse()

	switch mode {
	case "numbers":
		fmt.Printf("length: %s", length)
		fmt.Printf("min: %s", min)
		fmt.Printf("max: %s", max)
		var tmp *onetimecode.NumberedCode
		if length != -1 {
			tmp = onetimecode.NewNumberedCode(onetimecode.WithLength(length))
		} else {
			if min != -1 && max != -1 {
				tmp = onetimecode.NewNumberedCode(onetimecode.WithMinMax(min, max))
			} else {
				tmp = onetimecode.NewNumberedCode()
			}
		}
		code = tmp.String()
	case "alphanum":
		code = onetimecode.AlphaNumberCode(length)
	case "alphanumuc":
		code = onetimecode.AlphaNumberUcCode(length)
	case "uuid":
		code = onetimecode.Uuid()
	}

	fmt.Println(code)
}
